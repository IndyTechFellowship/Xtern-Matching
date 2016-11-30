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

func Register(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	dat := make(map[string]interface{})
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		log.Println(err.Error())
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
	
	//Don't want to send back the passwords for security resaons
	for i := 0; i < len(users); i++ {
		users[i].Password = "********"
	}
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request){
	contextUser := context.Get(r, "user")
	token, _ := contextUser.(*jwt.Token)
	if token.Valid {
		mapClaims := token.Claims.(jwt.MapClaims)
		org := strings.TrimSpace(mapClaims["org"].(string))
		role := strings.TrimSpace(mapClaims["role"].(string))

		type Response struct {
			Org   string	`json:"organization"`
			Role string	`json:"role"`
		}

		res := Response{Org: org, Role: role}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

func PutUser(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	err := services.UpdateUser(ctx, &user)
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

func Login(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	token, err := services.Login(ctx, dat["email"], dat["password"])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	dat = make(map[string]interface{})
	dat["token"] = string(token)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dat)
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
