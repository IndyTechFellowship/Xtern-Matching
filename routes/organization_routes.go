package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
)

func GetOrganizationRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/company", handlers.GetOrganizations).Methods("GET").Name("GetOrganization")
	router.HandleFunc("/company", handlers.AddOrganization).Methods("POST").Name("AddOrganization")
	router.HandleFunc("/company/addStudent", handlers.AddStudentToOrganization).Methods("POST").Name("AddStudentToOrganization")
	router.HandleFunc("/company/removeStudent", handlers.RemoveStudentFromOrganization).Methods("DELETE").Name("RemoveStudentFromOrganization")
	router.HandleFunc("/company/moveStudent", handlers.MoveStudentInOrganization).Methods("PUT").Name("MoveStudentInOrganization")
	return router
}