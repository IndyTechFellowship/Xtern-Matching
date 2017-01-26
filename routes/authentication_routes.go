package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
	_ "github.com/someone1/gcp-jwt-go"
	//"net/http"
	//"google.golang.org/appengine"
	//"github.com/gorilla/context"
)

func GetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/auth/login", handlers.Login).Methods("POST").Name("Login")
	return router
}
