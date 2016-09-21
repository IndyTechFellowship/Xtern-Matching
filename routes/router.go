package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handlers.Index).Methods("GET").Name("Index")
	router = SetStudentRoutes(router)
	return router
}
