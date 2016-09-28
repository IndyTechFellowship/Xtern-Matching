package services

import (
	"google.golang.org/appengine/datastore"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/someone1/gcp-jwt-go"
	"golang.org/x/net/context"
	//"time"
	"errors"
	"Xtern-Matching/models"
)

func Login(ctx context.Context,user models.User) ([]byte, error) {
	//TODO IMPROVE EVENTUALLY
	q := datastore.NewQuery("User").Filter("Email =", user.Email)

	var account models.User
	for t := q.Run(ctx); ; {
		_, err := t.Next(&account)
		if err == datastore.Done {
			return []byte(""), errors.New("User doesn't exist")
		}
		break
	}
	if account.Email == user.Email && bcrypt.CompareHashAndPassword([]byte(account.Password),[]byte(user.Password)) == nil {
		// Create a new token object, specifying signing method and the claims
		// you would like it to contain.
		token := jwt.New(jwt.GetSigningMethod("AppEngine"))
		//claims := token.Claims.(jwt.MapClaims)
		////token.Claims["iat"] = time.Now().Unix()
		//token.Claims["exp"] = time.Now().Add(time.Hour * time.Duration(24)).Unix()

		tokenString, err := token.SignedString(ctx)
		if err != nil {
			return []byte(""), err
		}
		return []byte(tokenString), err
	}
	return []byte(""), errors.New("Wrong Password")
}