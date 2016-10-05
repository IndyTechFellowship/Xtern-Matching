package services

import (
	"golang.org/x/net/context"
	"Xtern-Matching/models"
	"net/http"
	"google.golang.org/appengine/datastore"
)

// func AddStudentIdToCompanyList(ctx context.Context,student models.Student) (int,error) {
// 	// key := datastore.NewIncompleteKey(ctx, "Student", nil)
// 	if _, err := datastore.Put(ctx, &student); err != nil {
// 		return http.StatusInternalServerError, err
// 	}
// 	return http.StatusAccepted, nil
// }

// func GetStudent(ctx context.Context,_id int64) (models.Student,error) {
// 	studentKey := datastore.NewKey(ctx, "Student", "", _id, nil)
// 	var student models.Student
// 	if err := datastore.Get(ctx, studentKey, &student); err != nil {
// 		return models.Student{}, err
// 	}
// 	return student, nil
// }

func NewCompany(ctx context.Context,company models.Company) (int,error) {
	key := datastore.NewIncompleteKey(ctx, "Company", nil)
	if _, err := datastore.Put(ctx, key, &company); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusAccepted, nil
}
