package routes

import "github.com/gorilla/mux"

func SetSalesforceRoutes(router *mux.Router) *mux.Router {
	/*router.HandleFunc("/user/{userId}/", handlers.GetUser).Methods("GET").Name("GetUser")
	router.HandleFunc("/user/", handlers.CreateUser).Methods("POST").Name("CreateUser")
	router.HandleFunc("/user/", handlers.GetUsers).Methods("GET").Name("GetUsers")*/
	return router
}
