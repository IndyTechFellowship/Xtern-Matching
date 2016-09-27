package handlers

import (
	"net/http"
	"google.golang.org/appengine"
	"models"
	"encoding/json"
	"handlers/services"
)

func Register(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	responseStatus, err := services.Register(ctx,user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(responseStatus)
}

