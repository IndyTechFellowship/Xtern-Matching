package handlers

import (
	"Xtern-Matching/handlers/services"
	"Xtern-Matching/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

func GetStudent(w http.ResponseWriter, r *http.Request) {
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

func GetStudents(w http.ResponseWriter, r *http.Request) {
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

func ExportStudents(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	students, err := services.ExportStudents(ctx)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/csv")
	w.Write(students)
}

func GetStudentsFromIds(w http.ResponseWriter, r *http.Request) {
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

func PostStudent(w http.ResponseWriter, r *http.Request) {
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

func AddComment(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	newComment := models.Comment{Author: dat["author_name"].(string), Group: dat["group_name"].(string), Text: dat["text"].(string)}
	studentId := int64(dat["id"].(float64))
	student, err := services.GetStudent(ctx, studentId)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}

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

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	commentToDelete := models.Comment{Author: dat["author_name"].(string), Group: dat["group_name"].(string), Text: dat["text"].(string)}
	studentId := int64(dat["id"].(float64))
	student, err := services.GetStudent(ctx, studentId)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}

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

//8 MB file limit
const MAX_MEMORY = 8 * 1024 * 1024

func PostPDF(w http.ResponseWriter, r *http.Request) {

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
		num_id, _ := strconv.ParseInt(id, 10, 64)
		err := services.UpdateResume(ctx, num_id, file)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	}
}
