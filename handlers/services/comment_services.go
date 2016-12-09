package services

import (
	"Xtern-Matching/models"
	"net/http"
	"google.golang.org/appengine/datastore"
	"golang.org/x/net/context"
)

func GetComments(ctx context.Context, studentKey *datastore.Key, organizationKey *datastore.Key) ([]models.Comment,[]*datastore.Key,error) {
	q := datastore.NewQuery("Comment").Ancestor(studentKey)
	var allComments []models.Comment
	commentKeys, err := q.GetAll(ctx, &allComments)
	if err != nil {
		return nil, nil, err
	}

	//TODO optimize
	var keys []*datastore.Key
	comments := make([]models.Comment, 0)
	for index, comment := range allComments {
		if comment.Author.Parent().Equal(organizationKey) {
			comments = append(comments, comment)
			keys = append(keys, commentKeys[index])
		}
	}

	return comments, keys, nil
}

func AddComment(ctx context.Context, studentKey *datastore.Key, message string, author *datastore.Key) (int, error) {
	//studentKey := datastore.NewKey(ctx, "Student", "", studentId, nil)

	commentKey := datastore.NewIncompleteKey(ctx, "Comment", studentKey)
	comment := models.Comment{Message: message, Author: author}
	if _, err := datastore.Put(ctx, commentKey, &comment); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusAccepted, nil
}

func EditComment(ctx context.Context, commentKey *datastore.Key, message string) (int, error) {
	//studentKey := datastore.NewKey(ctx, "Student", "", studentId, nil)
	//commentKey := datastore.NewIncompleteKey(ctx, "Comment", &studentKey)

	var comment models.Comment
	err := datastore.Get(ctx, commentKey, &comment)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	comment.Message = message
	if _, err := datastore.Put(ctx, commentKey, &comment); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusAccepted, nil
}

func DeleteComment(ctx context.Context, commantKey *datastore.Key) (int, error) {
	if err := datastore.Delete(ctx, commantKey); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusAccepted, nil
}
