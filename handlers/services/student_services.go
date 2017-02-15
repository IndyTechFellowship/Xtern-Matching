package services

import (
	"Xtern-Matching/models"
	"archive/zip"
	"io"
	"net/http"
	"Xtern-Matching/handlers/services/csv"
	"os"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"bytes"
	"google.golang.org/appengine/urlfetch"
)

func StudentToStudentDecision(students []models.Student, keys []*datastore.Key) []models.StudentDecision {
	var studentDecisions []models.StudentDecision = make([]models.StudentDecision, len(students))
	for i := 0; i < len(students); i++ {
		student := students[i]
		studentDecisions[i] = models.StudentDecision{keys[i].IntID(), student.FirstName, student.LastName,
			student.GradYear, student.Grade, student.Gender, student.WorkStatus, student.Ethnicity, student.ReviewerGrades}
	}
	return studentDecisions
}

func GetStudentDecisionList(ctx context.Context, parent *datastore.Key) ([]models.StudentDecision, error) {

	students, keys, err := GetStudents(ctx, parent)
	if err != nil {
		return nil, err
	}
	return StudentToStudentDecision(students, keys), nil
}

func GetStudentsAtLeastWithStatus(ctx context.Context, status string) ([]models.StudentDecision, error) {
	statuses := [...]string{"Rejected (Stage 1)", "Rejected (Stage 2)", "Rejected (Stage 3)",
		"Undecided", "Stage 1 Approved", "Stage 2 Approved", "Stage 3 Approved"}
	query := datastore.NewQuery("Student")
	var students []models.StudentDecision
	for i := 0; i < len(statuses); i++ {

		if statuses[i] == status {
			for ; i < len(statuses); i++ {
				/*
					Not efficient query wise, but the best option for now
				 */
				newQuery := query.Filter("Status =", statuses[i])
				var newStudents []models.Student
				keys, err := newQuery.GetAll(ctx, &newStudents)
				if err != nil {
					return nil, err
				}
				newDecisions := StudentToStudentDecision(newStudents, keys)
				if newStudents != nil {
					for j := 0; j < len(newDecisions); j++ {
						students = append(students, newDecisions[j])
					}
				}
			}
			break
		}
	}
	return students, nil
}

func MoveStudentsToStatus(ctx context.Context, keys []*datastore.Key, status string) (error) {
	statuses := [...]string{"Rejected (Stage 1)", "Rejected (Stage 2)", "Rejected (Stage 3)",
		"Undecided", "Stage 1 Approved", "Stage 2 Approved", "Stage 3 Approved"}
	query := datastore.NewQuery("Student")

	for i := 0; i < len(statuses); i++ {
		if statuses[i] == status {
			break
		}
		newQuery := query.Filter("Status =", statuses[i])
		/*
			Terribly inefficient, but again,
			go datastore doesn't support the queries I would like to perform
 		*/
		var students []models.Student
		qKeys, err := newQuery.GetAll(ctx, &students)
		if err != nil {
			return err
		}
		if students != nil {
			for i := 0; i < len(students); i++ {
				for j := 0; j < len(keys); j++ {
					if keys[j].Equal(qKeys[i]) {
						students[i].Status = status
						datastore.Put(ctx, qKeys[i], &students[i])
					}
				}
			}
		}
	}
	return nil
}

func GetReviewedStudents(ctx context.Context, parent *datastore.Key) ([]models.StudentDecision, error) {
	students, err := GetStudentDecisionList(ctx, parent)
	if err != nil {
		return nil, err
	}
	var filteredStudents []models.StudentDecision
	for i := 0; i < len(students); i++ {
		if len(students[i].ReviewerGrades) > 0 {
			filteredStudents = append(filteredStudents, students[i])
		}
	}
	return filteredStudents, nil
}


/*
Gets all students in the database.
*/
func GetStudents(ctx context.Context, parent *datastore.Key) ([]models.Student, []*datastore.Key, error) {
	q := datastore.NewQuery("Student")
	if parent != nil {
		q = datastore.NewQuery("Student")
	}
	var students []models.Student
	keys, err := q.GetAll(ctx, &students)
	if err != nil {
		return nil, nil, err
	}

	return students, keys, nil
}

/*
Exports all student resumes in the Database.
Queries all students and exports them
 */
func ExportAllResumes(ctx context.Context) (*bytes.Buffer, error) {
	students, _, err := GetStudents(ctx, nil)
	if err != nil {
		return nil, err
	}
	return ExportResumes(ctx, students)
}

/*
Exports a slice of students as archive.pdf.
Useful for testing service to minimize the number of pdf GET requests
 */
func ExportResumes(ctx context.Context, students []models.Student) (*bytes.Buffer, error) {

	client := urlfetch.Client(ctx)
	buf := new(bytes.Buffer)

	archive := zip.NewWriter(buf)
	defer archive.Close()
	for _, student := range students {
		// Get the resume and write it
		resp, err := client.Get(student.Resume)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		f, err := archive.Create(student.Email + ".pdf")
		io.Copy(f, resp.Body)
	}

	return buf, nil
}

func GetStudent(ctx context.Context, studentKey *datastore.Key) (models.Student, error) {
	var student models.Student
	err := datastore.Get(ctx, studentKey, &student)
	if err != nil {
		return models.Student{}, err
	}
	return student, nil
}

func ExportStudents(ctx context.Context) ([]byte, error) {
	students, _, err := GetStudents(ctx, nil)
	if err != nil {
		return nil, err
	}
	output, err := csv.Marshal(students)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func NewStudent(ctx context.Context, student models.Student) (int, error) {
	key := datastore.NewIncompleteKey(ctx, "Student", nil)
	student.Active = true
	student.ReviewerGrades = make([]models.ReviewerGrade, 0, 1)

	//TODO make this done in a single put
	key, err := datastore.Put(ctx, key, &student)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	file, err := os.Open("public/sample.pdf")
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer file.Close()

	/* resumeURL, err := addResume(ctx, key.IntID(), file)
	if err != nil {
		log.Println("Error uploading resume")
		return http.StatusInternalServerError, err
	} */
	student.Resume = "http://localhost:8080/public/sample.pdf"//resumeURL

	_, err = datastore.Put(ctx, key, &student)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}

func SetStatus(ctx context.Context, studentKey *datastore.Key, status string) error {
	var student models.Student
	err := datastore.Get(ctx, studentKey, &student)
	if err != nil {
		return err
	}
	student.Status = status
	_, err = datastore.Put(ctx, studentKey, &student)
	if err != nil {
		return err
	}
	return nil
}

func SetGrade(ctx context.Context, studentKey *datastore.Key, grade float64) error {
	var student models.Student
	err := datastore.Get(ctx, studentKey, &student)
	if err != nil {
		return err
	}
	student.Grade = grade
	_, err = datastore.Put(ctx, studentKey, &student)
	if err != nil {
		return err
	}
	return nil
}