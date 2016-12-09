package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
)

func GetAdminRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/user", handlers.GetUsers).Methods("GET").Name("GetUsers")
	router.HandleFunc("/user/{userKey}", handlers.GetUser).Methods("GET").Name("GetUser")
	router.HandleFunc("/user", handlers.AddUser).Methods("POST").Name("AddUser")
	router.HandleFunc("/user/{userKey}", handlers.EditUser).Methods("PUT").Name("EditUser")
	router.HandleFunc("/user/{userKey}", handlers.DeleteUser).Methods("DELETE").Name("DeleteUser")
	return router
}