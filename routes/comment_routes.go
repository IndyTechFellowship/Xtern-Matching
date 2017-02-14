package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
)

func GetCommentRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/comment/{studentKey}", handlers.GetComments).Methods("GET").Name("GetComments")
	router.HandleFunc("/comment/{studentKey}", handlers.AddComment).Methods("POST").Name("AddComment")
	router.HandleFunc("/comment/{commentKey}", handlers.EditComment).Methods("PUT").Name("EditComment")
	router.HandleFunc("/comment/{commentKey}", handlers.DeleteComment).Methods("DELETE").Name("DeleteComment")
	return router
}
