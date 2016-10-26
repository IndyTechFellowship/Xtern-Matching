package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
)

func GetStudentRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/student/{Id}", handlers.GetStudent).Methods("GET").Name("GetStudent")
	router.HandleFunc("/student", handlers.GetStudents).Methods("GET").Name("GetStudents")
	router.HandleFunc("/student/resume/{Id}", handlers.PostPDF).Methods("POST").Name("CreatePDF")
	router.HandleFunc("/student", handlers.PostStudent).Methods("POST").Name("CreateStudent")
	router.HandleFunc("/student/getstudents", handlers.GetStudentsFromIds).Methods("POST").Name("GetStudentsFromIds")
	router.HandleFunc("/student/addComment", handlers.AddComment).Methods("POST").Name("AddComment")
	router.HandleFunc("/student/deleteComment", handlers.DeleteComment).Methods("POST").Name("DeleteComment")
	return router
}