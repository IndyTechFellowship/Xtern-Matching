package services

import (
	"Xtern-Matching/models"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	storage "google.golang.org/api/storage/v1"
	"google.golang.org/appengine/datastore"
)

func NewStudent(ctx context.Context, student models.Student) (int, error) {
	if reflect.DeepEqual(student, (models.Student{})) {
		return http.StatusBadRequest, errors.New("Not a proper student")
	}

	key := datastore.NewIncompleteKey(ctx, "Student", nil)
	student.Resume = "public/data_mocks/sample.pdf"
	student.Active = true

	if _, err := datastore.Put(ctx, key, &student); err != nil {
		return http.StatusInternalServerError, err
	}
	student.Id = key.IntID()
	UpdateStudent(ctx, student)
	return http.StatusAccepted, nil
}

func GetStudent(ctx context.Context, _id int64) (models.Student, error) {
	studentKey := datastore.NewKey(ctx, "Student", "", _id, nil)
	var student models.Student
	err := datastore.Get(ctx, studentKey, &student)
	if err != nil {
		return models.Student{}, err
	}
	student.Id = _id
	return student, nil
}

func GetStudents(ctx context.Context) ([]models.Student, error) {
	q := datastore.NewQuery("Student")
	log.Printf("%v",q)
	var students []models.Student
	keys, err := q.GetAll(ctx, &students)
	if err != nil {
		return nil, err
	}
	log.Printf("%v",keys)

	return students, nil
}

func GetStudentsFromIds(ctx context.Context, _ids []int64) ([]models.Student, error) {
	studentKeys := make([]*datastore.Key, len(_ids))
	for i := 0; i < len(_ids); i++ {
		studentKeys[i] = datastore.NewKey(ctx, "Student", "", _ids[i], nil)
	}
	students := make([]models.Student, len(_ids))

	for i := 0; i < len(_ids); i++ {
		if err := datastore.Get(ctx, studentKeys[i], &students[i]); err != nil {
			return nil, err
		}
	}

	for i := 0; i < len(students); i++ {
		students[i].Id = studentKeys[i].IntID()
	}
	return students, nil
}

func UpdateResume(ctx context.Context, id int64, file io.Reader) error {
	sid := strconv.Itoa(int(id))
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
		return err
	}
	service, err := storage.New(client)
	if err != nil {
		return err
	}

	//Access Bucket and see if it exists
	if _, err := service.Buckets.Get(bucketName).Do(); err == nil {
		log.Printf("Bucket %s already exists - skipping buckets.insert call.", bucketName)
	} else {
		// Create a bucket.
		if res, err := service.Buckets.Insert(projectID, &storage.Bucket{Name: bucketName}).Do(); err == nil {
			log.Printf("Created bucket %v at location %v\n\n", res.Name, res.SelfLink)
		} else {
			return err
		}
	}

	//Delete old resume copy if it exists
	if err := service.Objects.Delete(bucketName, sid+".pdf").Do(); err != nil {
		// If the object exists but wasn't deleted, the bucket deletion will also fail.
		log.Printf("Could not delete object during cleanup: %v\n\n", err)
	} else {
		log.Printf("Successfully deleted %s/%s during cleanup.\n\n", bucketName, sid)
	}

	//Insert new resume copy
	object := &storage.Object{Name: sid + ".pdf"}
	res, err := service.Objects.Insert(bucketName, object).Media(file).Do()
	if err == nil {
		log.Printf("Created object %v at location %v\n\n", res.Name, res.SelfLink)
	} else {
		return err
	}

	//Update student record to include resume link
	student, err := GetStudent(ctx, id)
	if err != nil {
		log.Println("Here")
		return err
	}
	student.Resume = res.MediaLink

	err = UpdateStudent(ctx, student)
	if err != nil {
		log.Println("Here1")
		return err
	}
	return nil
}
