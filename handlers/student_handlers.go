package handlers

import (
	"Xtern-Matching/handlers/services"
	"Xtern-Matching/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/appengine/datastore"

	"google.golang.org/appengine"
)

func GetStudents(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	students, keys, err := services.GetStudents(ctx, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	type Response struct {
		Keys     []*datastore.Key `json:"keys"`
		Students []models.Student `json:"students"`
	}
	response := Response{Keys: keys, Students: students}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	studentKey, err := datastore.DecodeKey(mux.Vars(r)["studentKey"])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	student, err := services.GetStudent(ctx, studentKey)

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

	status, err := services.NewStudent(ctx, student)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(status)
}

func SetGrade(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	studentKey, err := datastore.DecodeKey(mux.Vars(r)["studentKey"])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	grade := dat["grade"].(float64)

	err = services.SetGrade(ctx, studentKey, grade)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func SetStatus(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	studentKey, err := datastore.DecodeKey(mux.Vars(r)["studentKey"])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	status := dat["status"].(string)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = services.SetStatus(ctx, studentKey, status)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
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

func ExportResumes(w http.ResponseWriter,r *http.Request)  {
	ctx := appengine.NewContext(r)
	defer ctx.Done()

	buf, err := services.ExportAllResumes(ctx)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "archive/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=archive.zip")
	w.Write(buf.Bytes())
}
