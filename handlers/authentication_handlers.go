package handlers

import (
	"net/http"
	"google.golang.org/appengine"
	"encoding/json"
	"Xtern-Matching/handlers/services"
	"Xtern-Matching/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	token, err := services.Login(ctx,user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	type Message struct {
		Token string  `json:"token"`
	}
	var m Message
	m.Token = string(token)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	//ctx := appengine.NewContext(r)

	//var user models.User
	//decoder := json.NewDecoder(r.Body)
	//if err := decoder.Decode(&user); err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
	//responseStatus, err := services.Register(ctx,user)
	//if err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
	//w.WriteHeader(responseStatus)
}

