package services

import (
	"Xtern-Matching/models"
	"net/http"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"golang.org/x/crypto/bcrypt"
	"errors"
)

func Register(ctx context.Context,user models.User) (int,error) {
	count, err := datastore.NewQuery("User").Filter("Email =", user.Email).Count(ctx)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if count != 0 {
		//Successful Response, but user already exists
		// Should Update the user????
		return http.StatusAccepted, errors.New("User already exist")
	} else {
		key := datastore.NewIncompleteKey(ctx, "User", nil)
		
		//Hash Password
		pass, err := bcrypt.GenerateFromPassword([]byte(user.Password),14);
		if err != nil {
			return http.StatusInternalServerError, err
		}
		user.Password = string(pass)
		
		if _, err := datastore.Put(ctx, key, &user); err != nil {
			return http.StatusInternalServerError, err
		}
		return http.StatusCreated, nil
	}
}

func GetUsers(ctx context.Context, org string, role string) ([]models.User, error){
	query := datastore.NewQuery("User").Filter("Role =", role).Filter("Organization =", org)
	var users []models.User
	
	_, err := query.GetAll(ctx, &users)
	if err != nil {
		return nil, err
	}
	
	for i := 0; i < len(users); i++ {
		//users[i].Id = keys[i].IntID()
		users[i].Password = "********"
	}
	return users, err
}

func UpdateUser(ctx context.Context, user models.User) error {
	userKey := datastore.NewKey(ctx, "User", "", user.Id, nil)
	_,err := datastore.Put(ctx, userKey, &user)
	return err
}

func DeleteUser(ctx context.Context, id int64) error {
	userKey := datastore.NewKey(ctx, "User", "", id, nil)
	err := datastore.Delete(ctx, userKey)
	return err
}
