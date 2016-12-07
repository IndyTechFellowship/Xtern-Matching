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
	return http.StatusAccepted, nil
}

func UpdateStudent(ctx context.Context, student models.Student) error {
	studentKey := datastore.NewKey(ctx, "Student", "", student.Id, nil)
	_, err := datastore.Put(ctx, studentKey, &student)
	return err
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
	//log.Printf("%v",q)
	var students []models.Student
	keys, err := q.GetAll(ctx, &students)
	if err != nil {
		return nil, err
	}
	//log.Printf("%v",keys)

	for i := 0; i < len(students); i++ {
		students[i].Id = keys[i].IntID()
	}
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

func AddComment(ctx context.Context, studentId int64, newComment models.Comment) (int, error) {
	studentKey := datastore.NewKey(ctx, "Student", "", studentId, nil)
	var student models.Student

	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		err := datastore.Get(ctx, studentKey, &student)
		if err != nil && err != datastore.ErrNoSuchEntity {
			return err
		}
		student.Comments = ExtendComments(student.Comments, newComment)
		_, err = datastore.Put(ctx, studentKey, &student)
		return err
	}, nil)

	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusAccepted, nil
}

func DeleteComment(ctx context.Context, studentId int64, commentToDelete models.Comment) (int, error) {
	studentKey := datastore.NewKey(ctx, "Student", "", studentId, nil)
	var student models.Student

	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		err := datastore.Get(ctx, studentKey, &student)
		if err != nil && err != datastore.ErrNoSuchEntity {
			return err
		}
		student.Comments = removeComment(student.Comments, commentToDelete)
		_, err = datastore.Put(ctx, studentKey, &student)
		return err
	}, nil)

	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusAccepted, nil
}

////////////////////////////////////////////////
/////////////// HELPER FUNCTIONS ///////////////
////////////////////////////////////////////////

func ExtendComments(commentSlice []models.Comment, newComment models.Comment) []models.Comment {
	sliceSize := len(commentSlice)
	if sliceSize == cap(commentSlice) { // Check slice size versus max slice size
		newSlice := make([]models.Comment, len(commentSlice), 2*len(commentSlice)+1)
		copy(newSlice, commentSlice)
		commentSlice = newSlice
	}
	commentSlice = commentSlice[0 : sliceSize+1] // Grow size by 1
	commentSlice[sliceSize] = newComment
	return commentSlice
}

func removeComment(commentSlice []models.Comment, commentToRemove models.Comment) []models.Comment {
	filteredCommentSlice := commentSlice[:0]
	for _, comment := range commentSlice {
		if comment != commentToRemove {
			filteredCommentSlice = append(filteredCommentSlice, comment)
		}
	}
	return filteredCommentSlice
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
