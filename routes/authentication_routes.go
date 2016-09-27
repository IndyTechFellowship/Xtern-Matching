package routes

import (
	"github.com/gorilla/mux"
	"handlers"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/auth/login", handlers.Login).Methods("POST").Name("Login")
	router.HandleFunc("/auth/logout", handlers.Logout).Methods("POST").Name("Logout")
	return router
}
