package services

import (
	"golang.org/x/net/context"
	"Xtern-Matching/models"
	"net/http"
	"google.golang.org/appengine/datastore"
	"log"
)

func NewOrganization(ctx context.Context,org *models.Organization) (int,error) {
	key := datastore.NewIncompleteKey(ctx, "Organization", nil)
	if _, err := datastore.Put(ctx, key, org); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusAccepted, nil
}

func GetOrganization(ctx context.Context, _id int64) (models.Organization,error) {
	orgKey := datastore.NewKey(ctx, "Organization", "", _id, nil)
	var org models.Organization
	if err := datastore.Get(ctx, orgKey, &org); err != nil {
		return models.Organization{}, err
	}
	return org, nil
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

func MoveStudentInOrganization(ctx context.Context, orgKey datastore.Key, student1Id int64) (int64,error)  {
	//orgKey := datastore.NewKey(ctx, "Company", "", companyId, nil)
	var org models.Organization
	if err := datastore.Get(ctx, orgKey, &org); err != nil {
		return http.StatusInternalServerError, err
	}

	//TODO
	org.RemoveStudent(studentKey)

	if _, err := datastore.Put(ctx, orgKey, &org); err != nil {
		return http.StatusInternalServerError, err
	}
	return orgKey.IntID(), nil
}