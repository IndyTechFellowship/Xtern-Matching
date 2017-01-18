package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
)

func GetReviewerRoutes(router *mux.Router) *mux.Router {
	// router.HandleFunc("/organization", handlers.GetOrganizations).Methods("GET").Name("GetOrganization")
	router.HandleFunc("/reviewer/create", handlers.CreateReviewGroups).Methods("POST").Name("CreateReviewGroups")
	// router.HandleFunc("/organization/students", handlers.GetOrganizationStudents).Methods("GET").Name("GetOrganizationStudents")
	// router.HandleFunc("/organization/addStudent", handlers.AddStudentToOrganization).Methods("POST").Name("AddStudentToOrganization")
	// router.HandleFunc("/organization/removeStudent", handlers.RemoveStudentFromOrganization).Methods("POST").Name("RemoveStudentFromOrganization")
	// router.HandleFunc("/organization/moveStudent", handlers.MoveStudentInOrganization).Methods("PUT").Name("MoveStudentInOrganization")
	return router
}