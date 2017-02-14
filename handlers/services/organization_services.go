package services

import (
	"golang.org/x/net/context"
	"Xtern-Matching/models"
	"net/http"
	"google.golang.org/appengine/datastore"
)

func NewOrganization(ctx context.Context,name string) (*datastore.Key, error) {
	key := datastore.NewIncompleteKey(ctx, "Organization", nil)
	org := models.NewOrganization(name)
	key, err := datastore.Put(ctx, key, &org)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func GetOrganization(ctx context.Context, orgKey *datastore.Key) (models.Organization, error) {
	var org models.Organization
	err := datastore.Get(ctx, orgKey, &org)
	if err != nil {
		return models.Organization{}, err
	}
	return org, nil
}

func GetOrganizations(ctx context.Context) ([]models.Organization,[]*datastore.Key,error) {
	q := datastore.NewQuery("Organization").Project("Name")
	var orgs []models.Organization
	keys, err := q.GetAll(ctx, &orgs)
	if err != nil {
		return nil, nil, err
	}
	return orgs, keys, nil
}

func AddStudentToOrganization(ctx context.Context, orgKey *datastore.Key, studentKey *datastore.Key) (int64,error)  {
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

func RemoveStudentFromOrganization(ctx context.Context, orgKey *datastore.Key, studentKey *datastore.Key) error  {
	var org models.Organization
	if err := datastore.Get(ctx, orgKey, &org); err != nil {
		return err
	}

	org.RemoveStudent(studentKey)

	if _, err := datastore.Put(ctx, orgKey, &org); err != nil {
		return err
	}
	return nil
}

func MoveStudentInOrganization(ctx context.Context, orgKey *datastore.Key, studentKey *datastore.Key, pos int) (int64,error)  {
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