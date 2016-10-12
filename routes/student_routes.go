package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
)

func GetStudentRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/student/{Id}", handlers.GetStudent).Methods("GET").Name("GetStudent")
	router.HandleFunc("/student", handlers.GetStudents).Methods("GET").Name("GetStudents")
	router.HandleFunc("/student", handlers.PostStudent).Methods("POST").Name("CreateStudent")
	return router
}