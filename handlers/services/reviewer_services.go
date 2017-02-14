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

func UpdateReviewerGradeForStudent(ctx context.Context, reviewerKey *datastore.Key, studentKey *datastore.Key, reviewerGrade int) (error) {
	var student models.Student
	if err := datastore.Get(ctx, studentKey, &student); err != nil {
		return err
	}

	var doAppend bool = true

	for i := 0; i < len(student.ReviewerGrades); i++ {
		if student.ReviewerGrades[i].Reviewer == reviewerKey {
			student.ReviewerGrades[i].Grade = reviewerGrade
			doAppend = false
		}
	}
	if doAppend {
		var newGrade models.ReviewerGrade
		newGrade.Reviewer = reviewerKey
		newGrade.Grade = reviewerGrade
		student.ReviewerGrades = make([]models.ReviewerGrade, 0)
		student.ReviewerGrades = append(student.ReviewerGrades, newGrade)
	}

	if _, err := datastore.Put(ctx, studentKey, &student); err != nil {
		return err
	}
	return nil
}

func GetReviewerGradeForStudent(ctx context.Context, reviewerKey *datastore.Key, studentKey *datastore.Key) (int, error) {
	student, err := GetStudent(ctx, studentKey)
	if err != nil {
		return -1, errors.New("Could not find student")
	}

	for i := 0; i < len(student.ReviewerGrades); i++ {
		if student.ReviewerGrades[i].Reviewer.Equal(reviewerKey) {
			return int(student.ReviewerGrades[i].Grade), nil
		}
	}

	return -2, errors.New("Could not find grade for reviewer key")
}