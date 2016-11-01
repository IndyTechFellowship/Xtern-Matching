package handlers

import (
	"net/http"
	"google.golang.org/appengine"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"Xtern-Matching/models"
	"Xtern-Matching/handlers/services"
	"log"
)

func GetStudent(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	if id, ok := mux.Vars(r)["Id"]; ok {
		num_id, _ := strconv.ParseInt(id, 10, 64)
		student, err := services.GetStudent(ctx, num_id)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
		// if student.Resume == "" {
		// 	student.Resume = "public/data_mocks/sample.pdf"
		// }
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
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

// Add Student
func PostStudent(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var students []models.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&students); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	for _, student := range students {
		_, err := services.NewStudent(ctx, student)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
	}
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
}

//8 MB file limit
const MAX_MEMORY = 8 * 1024 * 1024
func PostPDF(w http.ResponseWriter,r *http.Request){

	//Get context and storage service
	ctx := appengine.NewContext(r)
	
	//Make sure pdf is less than 8 MB
	if err := r.ParseMultipartForm(MAX_MEMORY); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	
	//Fetch file from formdata
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	
	if id, ok := mux.Vars(r)["Id"]; ok {
		num_id, _ := strconv.ParseInt(id,10,64)
		err := services.UpdateResume(ctx, num_id, file)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	}
}
