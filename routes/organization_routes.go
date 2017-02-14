package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
)

func GetOrganizationRoutes(router *mux.Router) *mux.Router {
	// router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/organization", handlers.GetOrganizations).Methods("GET").Name("GetOrganization")
	router.HandleFunc("/organization/current", handlers.GetCurrentOrganization).Methods("GET").Name("GetCurrentOrganization")
	router.HandleFunc("/organization", handlers.AddOrganization).Methods("POST").Name("AddOrganization")
	router.HandleFunc("/organization/students", handlers.GetOrganizationStudents).Methods("GET").Name("GetOrganizationStudents")
	router.HandleFunc("/organization/addStudent", handlers.AddStudentToOrganization).Methods("POST").Name("AddStudentToOrganization")
	router.HandleFunc("/organization/removeStudent", handlers.RemoveStudentFromOrganization).Methods("POST").Name("RemoveStudentFromOrganization")
	// router.HandleFunc("/organization/moveStudent", handlers.MoveStudentInOrganization).Methods("PUT").Name("MoveStudentInOrganization")
	router.HandleFunc("/organization/switchStudents", handlers.SwitchStudentsInOrganization).Methods("PUT").Name("SwitchStudentsInOrganization")
	return router
}