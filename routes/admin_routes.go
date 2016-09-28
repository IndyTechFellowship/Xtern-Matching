package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
)

func SetAdminRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/admin/register", handlers.Register).Methods("POST").Name("Register")
	return router
}