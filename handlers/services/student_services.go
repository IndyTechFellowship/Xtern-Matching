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

func GetStudents(ctx context.Context, parent *datastore.Key) ([]models.Student, []*datastore.Key, error) {
	q := datastore.NewQuery("Student")
	if parent != nil {
		q = datastore.NewQuery("Student").Ancestor(parent)
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

func SetStatus(ctx context.Context, studentKey *datastore.Key, status string)  error {
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

func SetGrade(ctx context.Context, studentKey *datastore.Key, grade float64)  error {
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