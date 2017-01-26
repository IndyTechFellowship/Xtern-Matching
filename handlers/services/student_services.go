package services

import (
	"Xtern-Matching/models"
	"log"
	"net/http"

	"Xtern-Matching/handlers/services/csv"

	"io"
	"os"
	"strconv"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/storage/v1"
	"google.golang.org/appengine/datastore"
)

func GetStudentDecisionList(ctx context.Context, parent *datastore.Key) ([]models.StudentDecision, error) {
	q := datastore.NewQuery("Student")
	if parent != nil {
		q = datastore.NewQuery("Student").Ancestor(parent)
	}
	var students []models.StudentDecision
	_, err := q.GetAll(ctx, &students)
	if err != nil {
		return nil, err
	}
	return students, nil
}



func GetStudentsAtLeastStatus(ctx context.Context, status string) ([]models.StudentDecision, error) {
	statuses := [...]string{"Rejected (Stage 1)", "Rejected (Stage 2)", "Rejected (Stage 3)",
				"Undecided", "Stage 1 Approved", "Stage 2 Approved", "Stage 3 Approved"}
	query := datastore.NewQuery("Student")
	for i := 0; i < len(status); i++ {
		if statuses[i] == status {
			for ; i < len(status); i++ {
				query = query.Filter("status =", statuses[i])
			}
		}
	}
	var students []models.StudentDecision
	_, err := query.GetAll(ctx, &students)
	if err != nil {
		return nil, err
	}
	return students, nil
}

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

func GetStudent(ctx context.Context, studentKey *datastore.Key) (models.Student, error) {
	//studentKey := datastore.NewKey(ctx, "Student", "", _id, nil)
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
	resumeURL, err := addResume(ctx, key.IntID(), file)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	student.Resume = resumeURL
	_, err = datastore.Put(ctx, key, &student)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}

func addResume(ctx context.Context, studentId int64, file io.Reader) (string, error) {
	var bucketName string
	var projectID string
	if os.Getenv("XTERN_ENVIRONMENT") != "production" {
		bucketName = "xtern-matching-143216.appspot.com" //DEV Server
		projectID = "xtern-matching-143216"
	} else {
		bucketName = "xtern-matching.appspot.com"
		projectID = "xtern-matching"
	}

	client, err := google.DefaultClient(ctx, storage.DevstorageFullControlScope)
	if err != nil {
		return "", err
	}
	service, err := storage.New(client)
	if err != nil {
		return "", err
	}

	//Access Bucket and see if it exists
	if _, err := service.Buckets.Get(bucketName).Do(); err == nil {
		log.Printf("Bucket %s already exists - skipping buckets.insert call.", bucketName)
	} else {
		// Create a bucket.
		if res, err := service.Buckets.Insert(projectID, &storage.Bucket{Name: bucketName}).Do(); err == nil {
			log.Printf("Created bucket %v at location %v\n\n", res.Name, res.SelfLink)
		} else {
			return "", err
		}
	}

	//Insert new resume copy
	object := &storage.Object{Name: strconv.FormatInt(studentId, 10) + ".pdf"}
	res, err := service.Objects.Insert(bucketName, object).Media(file).Do()
	if err == nil {
		log.Printf("Created object %v at location %v\n\n", res.Name, res.SelfLink)
	} else {
		return "", err
	}

	return res.MediaLink, nil
}
