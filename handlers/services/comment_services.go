package services

import (
	"Xtern-Matching/models"
	"net/http"
	"google.golang.org/appengine/datastore"
	"golang.org/x/net/context"
	"log"
)

func GetComments(ctx context.Context, studentKey datastore.Key, organizationKey datastore.Key) ([]models.Comment, error) {
	q := datastore.NewQuery("Comment").Ancestor(studentKey)
	log.Printf("%v",q)
	var allComments []models.Comment
	keys, err := q.GetAll(ctx, &allComments)
	if err != nil {
		return nil, err
	}
	log.Printf("%v",keys)

	//TODO optimize
	comments := make([]datastore.Key, 0)
	for _, comment := range allComments {
		org := comment.Author.Parent()
		if org == organizationKey {
			comments = append(comments, comment)
		}
	}

	return comments, nil
}

func AddComment(ctx context.Context, studentKey datastore.Key, message string, author datastore.Key) (int, error) {
	//studentKey := datastore.NewKey(ctx, "Student", "", studentId, nil)

	commentKey := datastore.NewIncompleteKey(ctx, "Comment", &studentKey)
	comment := models.Comment{Message: message, Author: author}
	if _, err := datastore.Put(ctx, commentKey, &comment); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusAccepted, nil
}

func DeleteComment(ctx context.Context, commantKey datastore.Key) (int, error) {
	if err := datastore.Delete(ctx, commantKey); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusAccepted, nil
}
