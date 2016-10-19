package services

import (
	"golang.org/x/net/context"
	"Xtern-Matching/models"
	"net/http"
	"google.golang.org/appengine/datastore"
)

func NewStudent(ctx context.Context,student models.Student) (int,error) {
	key := datastore.NewIncompleteKey(ctx, "Student", nil)
	if _, err := datastore.Put(ctx, key, &student); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusAccepted, nil
}

func GetStudent(ctx context.Context,_id int64) (models.Student,error) {
	studentKey := datastore.NewKey(ctx, "Student", "", _id, nil)
	var student models.Student
	if err := datastore.Get(ctx, studentKey, &student); err != nil {
		return models.Student{}, err
	}
	student.Id = studentKey.IntID()
	return student, nil
}

func GetStudents(ctx context.Context) ([]models.Student,error) {
	q := datastore.NewQuery("Student")
	var students []models.Student
	keys, err := q.GetAll(ctx,&students)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(students); i++ {
		students[i].Id = keys[i].IntID()
	}
	return students, nil
}

func GetStudentsFromIds(ctx context.Context, _ids []int64) ([]models.Student,error) {
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

func AddComment(ctx context.Context, studentId int64, newComment models.Comment) (int,error) {
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