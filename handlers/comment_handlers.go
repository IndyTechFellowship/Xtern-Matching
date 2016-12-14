package handlers

import (
	"net/http"
	"encoding/json"
	"Xtern-Matching/handlers/services"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"github.com/gorilla/context"
	"github.com/dgrijalva/jwt-go"
	"Xtern-Matching/models"
	"github.com/gorilla/mux"
	"log"
)

func GetComments(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	studentKey, err := datastore.DecodeKey(dat["studentKey"].(string))
	if err != nil {
		log.Println(err.Error())
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

	comments, keys, err := services.GetComments(ctx, studentKey ,orgKey)
	if err != nil {
		//log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
	type Response struct {
		Keys []*datastore.Key		`json:"keys"`
		Comments []models.Comment	`json:"comments"`
	}
	response := Response{Keys: keys, Comments: comments}

	json.NewEncoder(w).Encode(response)
}

func AddComment(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	studentKey, err := datastore.DecodeKey(dat["studentKey"].(string))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	message :=  dat["message"].(string)
	user := context.Get(r, "user")
	mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	author, err := datastore.DecodeKey(mapClaims["key"].(string))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	status, err := services.AddComment(ctx, studentKey, message, mapClaims["name"].(string),author)
	if err != nil {
		//log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(status)
}

func EditComment(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	commentKey, err := datastore.DecodeKey(mux.Vars(r)["commentKey"])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	message := dat["message"].(string)
	status, err := services.EditComment(ctx, commentKey, message)
	if err != nil {
		//log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(status)
}

func DeleteComment(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	commentKey, err := datastore.DecodeKey(mux.Vars(r)["commentKey"])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	_, err = services.DeleteComment(ctx, commentKey)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}
