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