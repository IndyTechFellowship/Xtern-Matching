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

// func AddStudent(w http.ResponseWriter,r *http.Request) {
// 	ctx := appengine.NewContext(r)

// 	var studentIds []int64
// 	decoder := json.NewDecoder(r.Body)

// 	if err := decoder.Decode(&students); err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}

// 	for _, int64 := range studentIds {
// 		_, err := services.AddStudentIdToCompanyList(ctx, student)
// 		if err != nil {
// 			http.Error(w, err.Error(), 500)
// 			return
// 		}
// 	}
// 	w.WriteHeader(http.StatusOK)
// }

// func GetStudents(w http.ResponseWriter,r *http.Request) {
// 	ctx := appengine.NewContext(r)
// 	students, err := services.GetStudents(ctx)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}

// 	w.Header().Add("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(students)
// }

func PostCompany(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var companies []models.Company
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&companies); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	for _, company := range companies {
		_, err := services.NewCompany(ctx, company)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}

// 5733953138851840 

func GetCompany(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	if id, ok := mux.Vars(r)["Id"]; ok {
		num_id, _ := strconv.ParseInt(id, 10, 64)
		company, err := services.GetCompany(ctx, num_id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(company)
	}
	w.WriteHeader(http.StatusInternalServerError)
}