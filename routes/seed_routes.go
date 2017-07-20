package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
)

func GetSeedRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/seed/orgs", handlers.SeedOrgs).Methods("GET").Name("seedOrgs")
	router.HandleFunc("/seed/students", handlers.SeedStudents).Methods("GET").Name("seedOrgs")
	return router
}
