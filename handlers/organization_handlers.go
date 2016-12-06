package handlers

import (
	"net/http"
	"google.golang.org/appengine"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"Xtern-Matching/handlers/services"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"log"
	"strings"
	"github.com/gorilla/context"
	"appengine/datastore"
)

func GetOrganization(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	if id, ok := mux.Vars(r)["Id"]; ok {
		num_id, _ := strconv.ParseInt(id, 10, 64)
		company, err := services.GetOrganization(ctx, num_id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(company)
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func GetCurrentOrganization(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)
	user := context.Get(r, "user")
	mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)

	orgKey := strings.TrimSpace(mapClaims["org"].(datastore.Key))

	org, err := services.GetOrganization(ctx, orgKey)
	if err != nil {
		log.Print("ERROR GETTING COMPANY")
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(org)
}

func PostOrganization(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	name := dat["name"].(string)
	kind := dat["kind"].(string)

	_, err := services.NewOrganization(ctx, name, kind)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func AddStudentToOrganization(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	studentKey := dat["studentId"].(datastore.Key);

	// Get the company id from the token org and call the service with it
	user := context.Get(r, "user")
	mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	orgKey := strings.TrimSpace(mapClaims["org"].(datastore.Key))
	_, err := services.AddStudentToOrganization(ctx, orgKey, studentKey)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func RemoveStudentFromOrganization(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	// Get the student ID from the request data
	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	studentId :=  int64(dat["studentId"].(float64));

	// Get the company id from the token org and call the service with it
	user := context.Get(r, "user")
    token, err := user.(*jwt.Token)
    if token.Valid {
        mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
        org := strings.TrimSpace(mapClaims["org"].(string))
		company_num_id, er1 := strconv.ParseInt(org, 10, 64)
		if er1 != nil {
			log.Print("ERROR PARSING STRING TO INT64")
			log.Print(er1)
		}
		_, err := services.RemoveStudentIdFromCompanyList(ctx, company_num_id, studentId)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), 500)
			return
		}
	w.WriteHeader(http.StatusOK)
    } else {
        fmt.Println(err)
    }
}

func SwitchStudentsInOrganization(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	studentKey :=  dat["studentKey"].(datastore.Key);
	position :=  dat["position"].(datastore.Key);

	// Get the company id from the token org and call the service with it
	user := context.Get(r, "user")
        mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
        orgKey := strings.TrimSpace(mapClaims["org"].(datastore.Key))

	_, err := services.MoveStudentInOrganization(ctx, orgKey, studentKey, position)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)

}