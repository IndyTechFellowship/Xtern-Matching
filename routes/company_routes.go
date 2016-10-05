package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
)

func GetCompanyRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/company/{Id}", handlers.AddStudent).Methods("POST").Name("AddStudent")
	// router.HandleFunc("/company", handlers.GetStudents).Methods("GET").Name("GetStudents")
	router.HandleFunc("/company", handlers.PostCompany).Methods("POST").Name("PostCompany")
	return router
}