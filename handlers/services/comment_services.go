package services

import (
	"Xtern-Matching/models"
	"google.golang.org/appengine/datastore"
	"golang.org/x/net/context"
	"github.com/pkg/errors"
)

func GetComments(ctx context.Context, studentKey *datastore.Key, organizationKey *datastore.Key) ([]models.Comment,[]*datastore.Key,error) {
	q := datastore.NewQuery("Comment").Ancestor(studentKey)
	var allComments []models.Comment
	commentKeys, err := q.GetAll(ctx, &allComments)
	if err != nil {
		return nil, nil, err
	}

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

func AddComment(ctx context.Context, studentKey *datastore.Key, message string, name string, authorKey *datastore.Key) (models.Comment,*datastore.Key,error) {
	commentKey := datastore.NewIncompleteKey(ctx, "Comment", studentKey)
	comment := models.Comment{Message: message, Author: authorKey, AuthorName: name}
	key, err := datastore.Put(ctx, commentKey, &comment);
	if err != nil {
		return models.Comment{},nil,err
	}
	return comment,key,nil
}

func EditComment(ctx context.Context, commentKey *datastore.Key, message string) (models.Comment,error) {
	var comment models.Comment
	err := datastore.Get(ctx, commentKey, &comment)
	if err != nil {
		return models.Comment{}, err
	}
	comment.Message = message
	if _, err := datastore.Put(ctx, commentKey, &comment); err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}

func DeleteComment(ctx context.Context, commentKey *datastore.Key, author *datastore.Key) error {
	var comment models.Comment
	if err := datastore.Get(ctx, commentKey, &comment); err != nil {
		return err
	}
	if !comment.Author.Equal(author) {
		return errors.New("Can't delete comment that isn't yours")
	}

	if err := datastore.Delete(ctx, commentKey); err != nil {
		return err
	}
	return nil
}
