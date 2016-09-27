package services

import (
	"models"
	"net/http"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

func Register(ctx context.Context,user models.User) int {
	count, err := datastore.NewQuery("User").Filter("Email =", user.Email).Count(ctx)
	if err != nil {
		return http.StatusInternalServerError
	} else if count != 0 {
		return http.StatusInternalServerError
	} else {
		key := datastore.NewIncompleteKey(ctx, "User", nil)
		if _, err := datastore.Put(ctx, key, &user); err != nil {
			return http.StatusInternalServerError
		}
		return http.StatusAccepted
	}
}
