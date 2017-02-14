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
	"log"
)

func Login(ctx context.Context, email string, password string) ([]byte, error) {
	q := datastore.NewQuery("User").Filter("Email =", email)

	var account models.User
	accountKey, err := q.Run(ctx).Next(&account)
	if err == datastore.Done || accountKey == nil {
		return []byte(""), errors.New("User doesn't exist")
	}
	if account.Email == email && bcrypt.CompareHashAndPassword([]byte(account.Password),[]byte(password)) == nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS512,jwt.MapClaims {
			"iat": time.Now().Unix(),
			"exp": time.Now().Add(time.Hour * time.Duration(24)).Unix(),
			"org": accountKey.Parent().Encode(),
			"key": accountKey.Encode(),
			"name": account.Name,
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

func Register(ctx context.Context, organizationKey *datastore.Key, user models.User) (int, *datastore.Key,error) {
	count, err := datastore.NewQuery("User").Ancestor(organizationKey).Filter("Email =", user.Email).Count(ctx)
	if err != nil {
		log.Printf("Error querying organizations: bad key: %v\n", organizationKey)
		return http.StatusInternalServerError, nil, err
	} else if count != 0 {
		return http.StatusBadRequest, nil, errors.New("User already exist")
	} else {
		//Hash Password
		pass, err := bcrypt.GenerateFromPassword([]byte(user.Password),14)
		if err != nil {
			log.Println("Error hashing user password in registration")
			return http.StatusInternalServerError, nil, err
		}
		user.Password = string(pass)

		key := datastore.NewIncompleteKey(ctx, "User", organizationKey)
		if key, err = datastore.Put(ctx, key, &user); err != nil {
			log.Println("Error inserting user into Database")
			return http.StatusInternalServerError, nil, err
		}
		return http.StatusCreated, key, nil
	}
}

func GetUsers(ctx context.Context, org *datastore.Key) ([]models.User, []*datastore.Key, error) {
	query := datastore.NewQuery("User").Project("Name", "Email")
	if org != nil {
		query = query.Ancestor(org)
	}

	var users []models.User
	keys, err := query.GetAll(ctx, &users)
	if err != nil {
		return nil, nil, err
	}
	return users, keys, nil
}

func GetUsersByOrgName(ctx context.Context, orgName string) ([]models.User, []*datastore.Key, error) {
	var orgKey *datastore.Key
	query := datastore.NewQuery("Organization").Filter("Name =", orgName).KeysOnly()
	orgs, _ := query.GetAll(ctx, nil)
	orgKey = orgs[0]

	return GetUsers(ctx, orgKey)
}

func GetUser(ctx context.Context, userKey *datastore.Key) (models.User, error){
	var user models.User
	err := datastore.Get(ctx, userKey, &user)
	if err != nil {
		return models.User{}, err
	}
	return user, err
}

func EditUser(ctx context.Context, userKey *datastore.Key, name string, email string, password string) error {
	var user models.User
	err := datastore.Get(ctx, userKey, &user)
	if err != nil {
		return err
	}
	user.Name = name
	user.Email = email
	if password != "" {
		pass, err := bcrypt.GenerateFromPassword([]byte(password), 14);
		if err != nil {
			return err
		}
		user.Password = string(pass)
	}
	if _, err := datastore.Put(ctx, userKey, &user); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteUser(ctx context.Context, userKey *datastore.Key) error {
	err := datastore.Delete(ctx, userKey)
	return err
}