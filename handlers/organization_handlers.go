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

	_, err := services.NewOrganization(ctx, name)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetOrganizationStudents(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	user := context.Get(r, "user")
	mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	orgKey, err := datastore.DecodeKey(mapClaims["org"].(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}


	org, err := services.GetOrganization(ctx,orgKey)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	students := make([]models.Student,0)
	for _, key := range org.Students {
		student, err := services.GetStudent(ctx, key)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		students = append(students, student)
	}
	type Response struct {
		Keys []*datastore.Key		`json:"keys"`
		Students []models.Student	`json:"students"`
	}
	response := Response{Keys: org.Students, Students: students}


	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func AddStudentToOrganization(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	studentKey, err := datastore.DecodeKey(dat["studentKey"].(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	user := context.Get(r, "user")
	mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	orgKey, err := datastore.DecodeKey(mapClaims["org"].(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	_, err = services.AddStudentToOrganization(ctx, orgKey, studentKey)
	if err != nil {
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

	studentKey, err := datastore.DecodeKey(dat["studentKey"].(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	user := context.Get(r, "user")
	mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	orgKey, err := datastore.DecodeKey(mapClaims["org"].(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = services.RemoveStudentFromOrganization(ctx, orgKey, studentKey)
	if err != nil {
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
	studentKey, err := datastore.DecodeKey(dat["studentKey"].(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	user := context.Get(r, "user")
	mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	orgKey, err := datastore.DecodeKey(mapClaims["org"].(string))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	position :=  int(dat["position"].(float64));

	_, err = services.MoveStudentInOrganization(ctx, orgKey, studentKey, position)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)

}