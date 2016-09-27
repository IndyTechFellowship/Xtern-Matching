package services

import (
	"net/http"
	"google.golang.org/appengine/datastore"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/someone1/gcp-jwt-go"
	"golang.org/x/net/context"
	"models"
	"time"
)

func Login(ctx context.Context,user models.User) (int, []byte) {
	//TODO IMPROVE EVENTUALLY
	q := datastore.NewQuery("User").Filter("Email =", user.Email)

	for t := q.Run(ctx); ; {
		var account models.User
		_, err := t.Next(&user)
		if err == datastore.Done || err != nil {
			return http.StatusInternalServerError, []byte("")
		}
		if account.Email == user.Email && bcrypt.CompareHashAndPassword(user.Password, account.Password) == nil {
			// Create a new token object, specifying signing method and the claims
			// you would like it to contain.
			token := jwt.New(jwt.GetSigningMethod("AppEngine"))
			token.Claims["iat"] = time.Now().Unix()
			token.Claims["exp"] = time.Now().Add(time.Hour * time.Duration(24)).Unix()

			tokenString, err := token.SignedString(ctx)
			if err != nil {
				return http.StatusInternalServerError, []byte("")
			}
			return http.StatusAccepted, []byte(tokenString)
		}
	}
	return http.StatusUnauthorized, []byte("")
}