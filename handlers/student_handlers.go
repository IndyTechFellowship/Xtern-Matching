package handlers

import (
	"net/http"
	"google.golang.org/appengine"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"Xtern-Matching/models"
	"Xtern-Matching/handlers/services"
)

func GetStudent(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	if id, ok := mux.Vars(r)["Id"]; ok {
		num_id, _ := strconv.ParseInt(id, 10, 64)
		student, err := services.GetStudent(ctx, num_id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(student)
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func GetStudents(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)
	students, err := services.GetStudents(ctx)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func PostStudent(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var students []models.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&students); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	for _, student := range students {
		_, err := services.NewStudent(ctx, student)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
