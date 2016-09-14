package routes

import (
	"github.com/gorilla/mux"
	"github.com/Xtern-Matching/Backend/handlers"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	/*router.HandleFunc("/user/{userId}/", handlers.GetUser).Methods("GET").Name("GetUser")
	router.HandleFunc("/user/", handlers.CreateUser).Methods("POST").Name("CreateUser")
	router.HandleFunc("/user/", handlers.GetUsers).Methods("GET").Name("GetUsers")*/
	return router
}
