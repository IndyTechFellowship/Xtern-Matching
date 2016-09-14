package routes

import (
	"github.com/gorilla/mux"
	"bitbucket/LabyrinthAI/labyrinthback/handlers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handlers.Index).Methods("GET").Name("Index")
	router = SetSalesforceRoutes(router)
	return router
}
