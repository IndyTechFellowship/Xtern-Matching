package services

import (
	"Xtern-Matching/models"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"errors"
	"log"
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

func GetReviewGroupForReviewer(ctx context.Context, reviewerKey *datastore.Key) (models.ReviewGroup, *datastore.Key, error) {
	reviewGroups, reviewGroupKeys, err := GetReviewGroups(ctx, nil)
	if err != nil {
		return models.ReviewGroup{}, nil, err
	}

	for g := range reviewGroups {
		for r := range reviewGroups[g].Reviewers {
			log.Println(reviewGroups[g].Reviewers[r])
			if reviewerKey.Equal(reviewGroups[g].Reviewers[r]) {
				return reviewGroups[g], reviewGroupKeys[g], nil
			}
		}
	}

	return models.ReviewGroup{}, nil, errors.New("Could not find a ReviewGroup containing reviewer")
}