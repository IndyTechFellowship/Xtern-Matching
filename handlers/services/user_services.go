package services

import (
	"Xtern-Matching/models"
	"net/http"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"golang.org/x/crypto/bcrypt"
	"errors"
	"time"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/someone1/gcp-jwt-go"
)

func Register(ctx context.Context, organizationKey datastore.Key, user models.User) (int,error) {
	count, err := datastore.NewQuery("User").Ancestor(organizationKey).Count(ctx)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if count != 0 {
		return http.StatusBadRequest, errors.New("User already exist")
	} else {
		//Hash Password
		pass, err := bcrypt.GenerateFromPassword([]byte(user.Password),14);
		if err != nil {
			return http.StatusInternalServerError, err
		}
		user.Password = string(pass)

		key := datastore.NewIncompleteKey(ctx, "User", nil)
		if _, err := datastore.Put(ctx, key, &user); err != nil {
			return http.StatusInternalServerError, err
		}
		return http.StatusCreated, nil
	}
}

func GetUsers(ctx context.Context, org string, role string) ([]models.User, error){
	query := datastore.NewQuery("User").Filter("Role =", role).Filter("Organization =", org)
	var users []models.User
	
	keys, err := query.GetAll(ctx, &users)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(users); i++ {
		users[i].Id = keys[i].IntID()
	}
	return users, err
}

func GetUser(ctx context.Context, _id int64) (models.User, error){
	userKey := datastore.NewKey(ctx, "User", "", _id, nil)
	var user models.User
	err := datastore.Get(ctx, userKey, &user)
	if err != nil {
		return models.User{}, err
	}
	user.Id = userKey.IntID()
	return user, err
}

func UpdateUser(ctx context.Context, user *models.User) error {
	userKey := datastore.NewKey(ctx, "User", "", user.Id, nil)
	//In this case, the password wasn't updated, so fetch passsword from database and set it so it's not overwritten to ***
	if user.Password == "********" {
		oldUser, err := GetUser(ctx, user.Id)
		if err != nil {
			return err
		}
		user.Password = oldUser.Password
	} else {
		//Hash Password
		pass, err := bcrypt.GenerateFromPassword([]byte(user.Password),14);
		if err != nil {
			return err
		}
		user.Password = string(pass)
	}
	_,err := datastore.Put(ctx, userKey, user)
	return err
}

func DeleteUser(ctx context.Context, id int64) error {
	userKey := datastore.NewKey(ctx, "User", "", id, nil)
	err := datastore.Delete(ctx, userKey)
	return err
}

func Login(ctx context.Context, email string, password string) ([]byte, error) {
	q := datastore.NewQuery("User").Filter("Email =", email)


	var account models.User
	var accountKey datastore.Key
	for t := q.Run(ctx); ; {
		accountKey, err := t.Next(&account)
		if err == datastore.Done || accountKey == nil {
			return []byte(""), errors.New("User doesn't exist")
		}
		break
	}
	if account.Email == email && bcrypt.CompareHashAndPassword([]byte(account.Password),[]byte(password)) == nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS512,jwt.MapClaims {
			"iat": time.Now().Unix(),
			"exp": time.Now().Add(time.Hour * time.Duration(24)).Unix(),
			"org": accountKey.Parent(),
			"key": accountKey,
		})



		//TODO: Don't hardcode this here and in company_handlers.go
		tokenString, err := token.SignedString([]byte("My Secret"))
		if err != nil {
			return []byte(""), err
		}
		return []byte(tokenString), err
	}
	return []byte(""), errors.New("Wrong Password")
}