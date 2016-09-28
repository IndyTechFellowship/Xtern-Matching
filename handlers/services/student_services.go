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

func GetStudent(ctx context.Context,_id int) (models.Student,error) {
	studentKey := datastore.NewKey(ctx, "Student", "", _id, nil)
	var student models.Student
	if err := datastore.Get(ctx, studentKey, &student); err != nil {
		return nil, err
	}
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
