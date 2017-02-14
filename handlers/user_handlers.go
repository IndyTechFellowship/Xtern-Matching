package handlers

import (
	"net/http"
	"google.golang.org/appengine"
	"encoding/json"
	"Xtern-Matching/handlers/services"
	"Xtern-Matching/models"
	"log"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/appengine/datastore"
	"github.com/gorilla/mux"
)

func Login(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tokenString, err := services.Login(ctx, dat["email"].(string), dat["password"].(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	token, err := jwt.Parse(string(tokenString), func(token *jwt.Token) (interface{}, error) {
		return []byte("My Secret"), nil
	});
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	orgKey, err := datastore.DecodeKey(token.Claims.(jwt.MapClaims)["org"].(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	org, err := services.GetOrganization(ctx,orgKey)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	dat = make(map[string]interface{})
	dat["token"] = string(tokenString)
	dat["organizationName"] = org.Name


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dat)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var orgKey *datastore.Key
	if val, ok := mux.Vars(r)["orgKey"]; ok {
		// log.Println("Found orgkey");
		var err error
		orgKey, err = datastore.DecodeKey(val)
		if err != nil {
			log.Println("ERROR: ")
			http.Error(w, err.Error(), 500)
			return
		}
	} else {
		// log.Println("Grabbing all");
		orgKey = nil
	}
	// log.Printf("%v\n",orgKey)
	users, keys, err := services.GetUsers(ctx, orgKey)
	if err != nil {
		log.Println("ERROR: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	type Response struct {
		Keys []*datastore.Key		`json:"keys"`
		Users []models.User		`json:"users"`
	}
	response := Response{Keys: keys, Users: users}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetUsersByOrgName(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	if orgName, ok := mux.Vars(r)["orgName"]; ok {
		users, keys, err := services.GetUsersByOrgName(ctx, orgName)
		if err != nil {
			log.Println("ERROR: " + err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		type Response struct {
			Keys []*datastore.Key		`json:"keys"`
			Users []models.User		`json:"users"`
		}
		response := Response{Keys: keys, Users: users}

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "ERROR: could not find orgName", http.StatusInternalServerError)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	userKey, err := datastore.DecodeKey(mux.Vars(r)["userKey"])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

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

	orgKey, err := datastore.DecodeKey(dat["orgKey"].(string))
	if err != nil {
		log.Println("ERROR1");
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}


	responseStatus, err := services.Register(ctx, orgKey, user)
	if err != nil {
		log.Println("ERROR2");
		log.Println(err.Error())
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

	userKey, err := datastore.DecodeKey(mux.Vars(r)["userKey"])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	err = services.EditUser(ctx, userKey, name, email, password)
	if err != nil {
		log.Println("ERROR: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)

	userKey, err := datastore.DecodeKey(mux.Vars(r)["userKey"])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	err = services.DeleteUser(ctx, userKey)
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
