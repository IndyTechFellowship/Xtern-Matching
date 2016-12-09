package handlers

import (
	"net/http"
	"google.golang.org/appengine"
	"encoding/json"
	"Xtern-Matching/handlers/services"
	"github.com/dgrijalva/jwt-go"
	"log"
	"github.com/gorilla/context"
	"google.golang.org/appengine/datastore"
	"Xtern-Matching/models"
)

func GetOrganizations(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	organizations, keys, err := services.GetOrganizations(ctx)
	if err != nil {
		//log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
	type Response struct {
		Keys []*datastore.Key			`json:"keys"`
		Organizations []models.Organization	`json:"organizations"`
	}
	response := Response{Keys: keys, Organizations: organizations}
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func AddOrganization(w http.ResponseWriter,r *http.Request) {
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
	studentKey := dat["studentKey"].(*datastore.Key);

	user := context.Get(r, "user")
	mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	orgKey := mapClaims["org"].(*datastore.Key)

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

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	studentKey := dat["studentKey"].(*datastore.Key);

	user := context.Get(r, "user")
	mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	orgKey := mapClaims["org"].(*datastore.Key)

	_, err := services.RemoveStudentFromOrganization(ctx, orgKey, studentKey)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func MoveStudentInOrganization(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	studentKey :=  dat["studentKey"].(*datastore.Key);
	position :=  dat["position"].(int);

	user := context.Get(r, "user")
        mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
        orgKey := mapClaims["org"].(*datastore.Key)

	_, err := services.MoveStudentInOrganization(ctx, orgKey, studentKey, position)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)

}

//func GetOrganization(w http.ResponseWriter,r *http.Request) {
//	ctx := appengine.NewContext(r)
//
//	user := context.Get(r, "user")
//	mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
//	org := mapClaims["org"].(*datastore.Key)
//
//	org, err := services.GetOrganization(ctx, org)
//	if err != nil {
//		log.Print(err)
//		http.Error(w, err.Error(), 500)
//		return
//	}
//	w.Header().Add("Access-Control-Allow-Origin", "*")
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(org)
//}