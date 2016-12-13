package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
)

func GetCompanyRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/company/{Id}", handlers.GetCompany).Methods("GET").Name("GetCompany")
	router.HandleFunc("/company/addStudent", handlers.AddStudent).Methods("POST").Name("AddStudent")
	router.HandleFunc("/company/removeStudent", handlers.RemoveStudent).Methods("POST").Name("RemoveStudent")
	router.HandleFunc("/company/switchStudents", handlers.SwitchStudents).Methods("POST").Name("SwitchStudents")
	// router.HandleFunc("/company/getCurrentCompany", handlers.GetCurrentCompany).Methods("GET").Name("GetCurrentCompany") //broken?
	router.HandleFunc("/company/getCurrentCompany/{Id}", handlers.GetCurrentCompany).Methods("GET").Name("GetCurrentCompany") //use this one
	router.HandleFunc("/company", handlers.PostCompany).Methods("POST").Name("PostCompany")
	return router
}