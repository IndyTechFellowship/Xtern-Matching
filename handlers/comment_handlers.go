package handlers

import (
	"net/http"
	"encoding/json"
	"Xtern-Matching/handlers/services"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"github.com/gorilla/context"
	"github.com/dgrijalva/jwt-go"
)

func GetComments(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	studentKey := dat["studentKey"].(datastore.Key)
	user := context.Get(r, "user")
	mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	org := mapClaims["org"].(datastore.Key)
	comments, err := services.GetComments(ctx, studentKey, org)
	if err != nil {
		//log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(status)
}

func AddComment(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	studentKey := dat["studentKey"].(datastore.Key)
	message := dat["message"].(string)
	user := context.Get(r, "user")
	mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	author := mapClaims["key"].(datastore.Key)
	status, err := services.AddComment(ctx, studentKey, message, author)
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
	studentKey := dat["studentKey"].(datastore.Key)
	message := dat["message"].(string)
	status, err := services.EditComment(ctx, studentKey, message)
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
	commentKey := dat["key"].(datastore.Key)

	_, err := services.DeleteComment(ctx, commentKey)
	if err != nil {
		http.Error(w, err, 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}
