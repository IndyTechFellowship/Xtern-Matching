package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
)

func GetAdminRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/admin", handlers.GetUsers).Methods("GET").Name("GetUsers")
	router.HandleFunc("/admin/{userKey}", handlers.GetUser).Methods("GET").Name("GetUser")
	router.HandleFunc("/admin", handlers.AddUser).Methods("POST").Name("AddUser")
	router.HandleFunc("/admin/{userKey}", handlers.EditUser).Methods("PUT").Name("EditUser")
	router.HandleFunc("/admin/{userKey}", handlers.DeleteUser).Methods("DELETE").Name("DeleteUser")
	return router
}