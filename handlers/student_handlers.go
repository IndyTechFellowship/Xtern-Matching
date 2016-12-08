package handlers

import (
	"net/http"
	"google.golang.org/appengine"
	"encoding/json"
	"Xtern-Matching/models"
	"Xtern-Matching/handlers/services"
	"log"
	"github.com/gorilla/context"
	"google.golang.org/appengine/datastore"
)

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

func GetStudent(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	studentKey := context.Get(r, "studentKey")
	student, err := services.GetStudent(ctx, studentKey.(*datastore.Key))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func AddStudent(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var student models.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&student); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	//Make sure pdf is less than 8 MB
	//if err := r.ParseMultipartForm(8 * 1024 * 1024); err != nil {
	//	http.Error(w, err.Error(), http.StatusForbidden)
	//	return
	//}
	////Fetch file from formdata
	//file, _, err := r.FormFile("file")
	//if err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
	//TODO fix during form-stack implementation
	//file, err := os.Open("public/sample.pdf")
	//if err != nil {
	//	log.Println(err.Error())
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
	//defer file.Close()

	status,err := services.NewStudent(ctx, student)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(status)
}
