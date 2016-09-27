package routes

import (
	"github.com/gorilla/mux"
	"handlers"
)

func SetStudentRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/student/{Id}", handlers.GetStudent).Methods("GET").Name("GetStudent")
	router.HandleFunc("/student", handlers.GetStudent).Methods("GET").Name("GetStudents")
	router.HandleFunc("/student", handlers.PostStudent).Methods("POST").Name("CreateStudent")
	return router
}