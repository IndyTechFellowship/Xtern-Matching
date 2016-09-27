package handlers

import (
	"net/http"
	"handlers/services"
	"google.golang.org/appengine"
	"encoding/json"
	"models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	responseStatus, token := services.Login(ctx,user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	responseStatus := services.Register(ctx,user)
	w.WriteHeader(responseStatus)
}

