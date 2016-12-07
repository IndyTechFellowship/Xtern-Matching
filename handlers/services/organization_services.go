package services

import (
	"golang.org/x/net/context"
	"Xtern-Matching/models"
	"net/http"
	"google.golang.org/appengine/datastore"
	"log"
)

func NewOrganization(ctx context.Context,name string, kind string) (int,error) {
	key := datastore.NewIncompleteKey(ctx, "Organization", nil)
	org := models.NewOrganization(name, kind)
	if _, err := datastore.Put(ctx, key, &org); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusAccepted, nil
}

func GetOrganizations(ctx context.Context) ([]models.Organization,error) {
	q := datastore.NewQuery("Organization")
	log.Printf("%v",q)

	var orgs []models.Organization
	keys, err := q.GetAll(ctx, &orgs)
	if err != nil {
		return nil, err
	}

	log.Printf("%v", keys)
	return orgs, nil
}

func AddStudentToOrganization(ctx context.Context, orgKey datastore.Key, studentKey datastore.Key) (int64,error)  {
	//orgKey := datastore.NewKey(ctx, "Company", "", companyId, nil)
	var org models.Organization
	if err := datastore.Get(ctx, orgKey, &org); err != nil {
		return http.StatusInternalServerError, err
	}

	org.AddStudent(studentKey)

	if _, err := datastore.Put(ctx, orgKey, &org); err != nil {
		return http.StatusInternalServerError, err
	}
	return orgKey.IntID(), nil
}

func RemoveStudentFromOrganization(ctx context.Context, orgKey datastore.Key, studentKey datastore.Key) (int64,error)  {
	//orgKey := datastore.NewKey(ctx, "Company", "", companyId, nil)
	var org models.Organization
	if err := datastore.Get(ctx, orgKey, &org); err != nil {
		return http.StatusInternalServerError, err
	}

	org.RemoveStudent(studentKey)

	if _, err := datastore.Put(ctx, orgKey, &org); err != nil {
		return http.StatusInternalServerError, err
	}
	return orgKey.IntID(), nil
}

func MoveStudentInOrganization(ctx context.Context, orgKey datastore.Key, studentKey datastore.Key, pos int) (int64,error)  {
	//orgKey := datastore.NewKey(ctx, "Company", "", companyId, nil)
	var org models.Organization
	if err := datastore.Get(ctx, orgKey, &org); err != nil {
		return http.StatusInternalServerError, err
	}

	org.MoveStudent(studentKey, pos)

	if _, err := datastore.Put(ctx, orgKey, &org); err != nil {
		return http.StatusInternalServerError, err
	}
	return orgKey.IntID(), nil
}

//func GetOrganization(ctx context.Context, orgKey datastore.Key) (models.Organization,error) {
//	//orgKey := datastore.NewKey(ctx, "Organization", "", _id, nil)
//	var org models.Organization
//	if err := datastore.Get(ctx, orgKey, &org); err != nil {
//		return models.Organization{}, err
//	}
//	return org, nil
//}