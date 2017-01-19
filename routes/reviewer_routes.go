package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
)

func GetReviewerRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/reviewer/getReviewGroups", handlers.GetReviewGroups).Methods("GET").Name("GetReviewGroups")
	router.HandleFunc("/reviewer/create", handlers.CreateReviewGroups).Methods("POST").Name("CreateReviewGroups")
	router.HandleFunc("/reviewer/getReviewGroupForReviewer", handlers.GetReviewGroupForReviewer).Methods("POST").Name("GetReviewGroupForReviewer")
	// router.HandleFunc("/organization/addStudent", handlers.AddStudentToOrganization).Methods("POST").Name("AddStudentToOrganization")
	// router.HandleFunc("/organization/removeStudent", handlers.RemoveStudentFromOrganization).Methods("POST").Name("RemoveStudentFromOrganization")
	// router.HandleFunc("/organization/moveStudent", handlers.MoveStudentInOrganization).Methods("PUT").Name("MoveStudentInOrganization")


	return router
}