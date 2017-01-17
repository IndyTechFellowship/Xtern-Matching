package routes

import (
	"Xtern-Matching/handlers"

	"github.com/gorilla/mux"
)

func GetStudentRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/student/resumes", handlers.ExportResumes).Methods("GET").Name("ExportResumes")
	router.HandleFunc("/student/export", handlers.ExportStudents).Methods("GET").Name("ExportStudents")
	router.HandleFunc("/student/{studentKey}", handlers.GetStudent).Methods("GET").Name("GetStudent")
	router.HandleFunc("/student", handlers.GetStudents).Methods("GET").Name("GetStudents")
	router.HandleFunc("/student", handlers.AddStudent).Methods("POST").Name("AddStudent")
	return router
}
