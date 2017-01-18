package services

import (
	"Xtern-Matching/models"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

func GetReviewGroups(ctx context.Context, parent *datastore.Key) ([]models.ReviewGroup, []*datastore.Key, error) {
	q := datastore.NewQuery("ReviewGroup")
	if parent != nil {
		q = datastore.NewQuery("ReviewGroup").Ancestor(parent)
	}
	var reviewGroups []models.ReviewGroup
	keys, err := q.GetAll(ctx, &reviewGroups)
	if err != nil {
		return nil, nil, err
	}

	return reviewGroups, keys, nil
}