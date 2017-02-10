package services

import (
	"golang.org/x/net/context"
	"Xtern-Matching/models"
	"google.golang.org/appengine/datastore"
)

func AddForm(ctx context.Context, form models.Form) (*datastore.Key, error) {
	key := datastore.NewIncompleteKey(ctx, "Form", nil)
	key, err := datastore.Put(ctx, key, &form)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func GetActiveForm(ctx context.Context) (models.Form, *datastore.Key, error) {
	q := datastore.NewQuery("Form").Filter("Active=", true)
	var forms []models.Form
	keys, err := q.GetAll(ctx, &forms)
	if err != nil {
		return models.Form{}, nil, err
	}

	return forms[0], keys[0], nil
}