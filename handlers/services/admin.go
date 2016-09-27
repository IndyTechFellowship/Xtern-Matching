package services

import (
	"models"
	"net/http"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx context.Context,user models.User) (int,error) {
	count, err := datastore.NewQuery("User").Filter("Email =", user.Email).Count(ctx)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if count != 0 {
		return http.StatusInternalServerError, nil
	} else {
		key := datastore.NewIncompleteKey(ctx, "User", nil)
		pass, err := bcrypt.GenerateFromPassword([]byte(user.Password),14);
		if err != nil {
			return http.StatusInternalServerError, err
		}
		user.Password = string(pass)
		if _, err := datastore.Put(ctx, key, &user); err != nil {
			return http.StatusInternalServerError, err
		}
		return http.StatusAccepted, nil
	}
}
