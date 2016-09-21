package handlers

import (
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"Xtern-Matching/models"
	"encoding/json"
	"github.com/gorilla/mux"
)

func GetStudent(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	if _, ok := mux.Vars(r)["Id"]; ok {
		studentKey := datastore.NewKey(ctx, "Student", r.PostForm.Get("id"), 0, nil)
		var student models.Student
		if err := datastore.Get(ctx, studentKey, &student); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(student)
	} else {
		q := datastore.NewQuery("Student")
		var students []models.Student
		if _, err := q.GetAll(ctx,&students); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(students)
	}
}

func PostStudent(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var student models.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&student); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	key := datastore.NewIncompleteKey(ctx, "Student", nil)
	if _, err := datastore.Put(ctx, key, &student); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}
