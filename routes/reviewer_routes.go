package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
)

func GetReviewerRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/reviewer/getReviewGroups", handlers.GetReviewGroups).Methods("GET").Name("GetReviewGroups")
	router.HandleFunc("/reviewer/getCurrentProgress", handlers.GetCurrentProgress).Methods("GET").Name("GetCurrentProgress")
	router.HandleFunc("/reviewer/create", handlers.CreateReviewGroups).Methods("POST").Name("CreateReviewGroups")
	router.HandleFunc("/reviewer/getReviewGroupForReviewer", handlers.GetReviewGroupForReviewer).Methods("POST").Name("GetReviewGroupForReviewer")
	router.HandleFunc("/reviewer/getReviewerGradeForStudent/{studentKey}", handlers.GetReviewerGradeForStudent).Methods("GET").Name("GetReviewerGradeForStudent")
	router.HandleFunc("/reviewer/postReviewerGradeForStudent", handlers.PostReviewerGradeForStudent).Methods("POST").Name("PostReviewerGradeForStudent")

	return router
}