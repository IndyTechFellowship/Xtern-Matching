package handlers

import (
	"net/http"
	"google.golang.org/appengine"
	"github.com/gorilla/mux"
	"encoding/json"
	"Xtern-Matching/handlers/services"
	"Xtern-Matching/models"
	"log"
	"errors"	
	"strconv"
	"github.com/gorilla/context"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"google.golang.org/appengine/datastore"
)

func Login(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	token, err := services.Login(ctx, dat["email"].(string), dat["password"].(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	dat = make(map[string]interface{})
	dat["token"] = string(token)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dat)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	user := context.Get(r, "user")
	mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	org := mapClaims["org"].(datastore.Key)

	users, err := services.GetUsers(ctx, org)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	userKey := context.Get(r, "userKey").(datastore.Key)
	user, err := services.GetUser(ctx, userKey)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	dat := make(map[string]interface{})
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		//log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	var user models.User
	user.Name = dat["name"].(string)
	user.Email = dat["email"].(string)
	user.Password = dat["password"].(string)
	responseStatus, err := services.Register(ctx, dat["key"].(datastore.Key), user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(responseStatus)
}

func EditUser(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)

	dat := make(map[string]interface{})
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	name := dat["name"].(string)
	email := dat["email"].(string)
	password := dat["password"].(string)

	userKey := context.Get(r, "userKey").(datastore.Key)
	err := services.EditUser(ctx, userKey, name, email, password)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)

	userKey := context.Get(r, "userKey").(datastore.Key)
	err := services.DeleteUser(ctx, userKey)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

//func BulkRegister(w http.ResponseWriter, r *http.Request) {
//	ctx := appengine.NewContext(r)
//
//	var users []models.User
//	decoder := json.NewDecoder(r.Body)
//
//	if err := decoder.Decode(&users); err != nil {
//		log.Println(err.Error())
//		http.Error(w, err.Error(), 500)
//		return
//	}
//
//	var foundUser bool = false
//	var successfulUsers int = 0
//	var errorOccured bool = false
//
//	for _, user :=range users {
//		code, err := services.Register(ctx,user)
//		if err != nil {
//			if code == http.StatusAccepted {
//				//User already exists
//				foundUser = true
//			} else {
//				//Other error occured
//				log.Println(err.Error())
//				errorOccured = true
//			}
//		} else {
//			successfulUsers++
//		}
//	}
//
//	if !foundUser && successfulUsers > 0 && !errorOccured {
//		// All users added
//		http.Error(w, errors.New(fmt.Sprintf("Added %d new users.", successfulUsers)).Error(), http.StatusCreated)
//	} else if foundUser && successfulUsers > 0 && !errorOccured {
//		// only some for the users added
//		http.Error(w, errors.New(fmt.Sprintf("Some users already exist. Added %d new users.", successfulUsers)).Error(), http.StatusCreated)
//	} else if errorOccured && successfulUsers > 0 {
//		http.Error(w, errors.New(fmt.Sprintf("An Error occured. Added %d new users.", successfulUsers)).Error(), http.StatusCreated)
//	} else if foundUser && successfulUsers == 0 && !errorOccured {
//		http.Error(w, errors.New("All users already found. No new users added").Error(), http.StatusAccepted)
//	} else if !foundUser && successfulUsers == 0 && errorOccured {
//		http.Error(w, errors.New("An Error occured. No Users added").Error(), http.StatusInternalServerError)
//	} else{
//		http.Error(w, errors.New("An Error occured.").Error(), http.StatusInternalServerError)
//	}
//
//}
