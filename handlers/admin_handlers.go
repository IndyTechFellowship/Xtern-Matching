package handlers

import (
	"net/http"
	"google.golang.org/appengine"
	"github.com/gorilla/mux"
	"encoding/json"
	"Xtern-Matching/handlers/services"
	"Xtern-Matching/models"
	"log"
	"fmt"
	"errors"	
	"strconv"
)

func Register(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	responseStatus, err := services.Register(ctx,user)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(responseStatus)
}

func BulkRegister(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	
	var users []models.User
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&users); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	var foundUser bool = false
	var successfulUsers int = 0
	var errorOccured bool = false

	for _, user :=range users {
		code, err := services.Register(ctx,user)
		if err != nil {
			if code == http.StatusAccepted {
				//User already exists
				foundUser = true
			} else {
				//Other error occured
				log.Println(err.Error())
				errorOccured = true
			}
		} else {
			successfulUsers++
		}
	}

	if !foundUser && successfulUsers > 0 && !errorOccured {
		// All users added
		http.Error(w, errors.New(fmt.Sprintf("Added %d new users.", successfulUsers)).Error(), http.StatusCreated)
	} else if foundUser && successfulUsers > 0 && !errorOccured {
		// only some for the users added
		http.Error(w, errors.New(fmt.Sprintf("Some users already exist. Added %d new users.", successfulUsers)).Error(), http.StatusCreated)			
	} else if errorOccured && successfulUsers > 0 {
		http.Error(w, errors.New(fmt.Sprintf("An Error occured. Added %d new users.", successfulUsers)).Error(), http.StatusCreated)
	} else if foundUser && successfulUsers == 0 && !errorOccured {
		http.Error(w, errors.New("All users already found. No new users added").Error(), http.StatusAccepted)
	} else if !foundUser && successfulUsers == 0 && errorOccured {
		http.Error(w, errors.New("An Error occured. No Users added").Error(), http.StatusInternalServerError)
	} else{
		http.Error(w, errors.New("An Error occured.").Error(), http.StatusInternalServerError)
	}
	
}

func GetUsers(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)

	role, role_ok := mux.Vars(r)["Role"]
	org, org_ok := mux.Vars(r)["Organization"]
	
	if !role_ok || !org_ok {
		//params not found
		log.Println("Missing either Role or Organization")
		http.Error(w, errors.New("Missing either Role or Organization").Error(), http.StatusBadRequest)
		return
	}

	users, err := services.GetUsers(ctx, org, role)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	err := services.UpdateUser(ctx, user)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}	
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)
	id, id_ok := mux.Vars(r)["Id"]
	if !id_ok {
		log.Println("Missing Id to delete user")
		http.Error(w, errors.New("Missing Id to delete user").Error(), http.StatusBadRequest)
		return
	}
	num_id, _ := strconv.ParseInt(id, 10, 64)
	err := services.DeleteUser(ctx, num_id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
