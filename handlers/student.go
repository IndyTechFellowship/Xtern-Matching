package handlers

import (
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"Xtern-Matching/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func GetStudent(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	if id, ok := mux.Vars(r)["id"]; ok {
		num_id, _ := strconv.ParseInt(id,10,64)
		studentKey := datastore.NewKey(ctx, "Student", "", num_id, nil)
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
		keys, err := q.GetAll(ctx,&students)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		for i := 0; i < len(students); i++ {
			students[i].Id = keys[i].IntID()
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
