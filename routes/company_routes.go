package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
)

func GetCompanyRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/company/{Id}", handlers.GetCompany).Methods("GET").Name("GetCompany")
	// router.HandleFunc("/company", handlers.GetStudents).Methods("GET").Name("GetStudents")
	router.HandleFunc("/company", handlers.PostCompany).Methods("POST").Name("PostCompany")
	return router
}