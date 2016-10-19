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

func GetStudentsFromIds(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)
	type intIds struct {
		Id []int64 `json:"_ids"`
	}

	// var string_ids []string
	decoder := json.NewDecoder(r.Body)
	// if err := decoder.Decode(&string_ids); err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }

	var _ids intIds

	if err := decoder.Decode(&_ids); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}



	// var _ids []int64
	// for i := 0; i<len(string_ids); i++ {
	// 	id, _ := strconv.ParseInt(string_ids[i], 10, 64)
	// 	_ids[i] = id
	// }

	

	students, err := services.GetStudentsFromIds(ctx, _ids.Id)
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

func AddComment(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	if id, ok := mux.Vars(r)["Id"]; ok {
		num_id, _ := strconv.ParseInt(id, 10, 64)
		student, err := services.GetStudent(ctx, num_id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		newComment := models.Comment{Author: "golang test author", Group: "golang test org", Text: "golang test text"}

		_, addCommentErr := services.AddComment(ctx, student.Id, newComment)
		if addCommentErr != nil {
			http.Error(w, addCommentErr.Error(), 500)
			return
		}

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(student)
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func DeleteComment(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	if id, ok := mux.Vars(r)["Id"]; ok {
		num_id, _ := strconv.ParseInt(id, 10, 64)
		student, err := services.GetStudent(ctx, num_id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		commentToDelete := models.Comment{Author: "golang test author", Group: "golang test org", Text: "golang test text"}

		_, addCommentErr := services.DeleteComment(ctx, student.Id, commentToDelete)
		if addCommentErr != nil {
			http.Error(w, addCommentErr.Error(), 500)
			return
		}

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(student)
	}
	w.WriteHeader(http.StatusInternalServerError)
}