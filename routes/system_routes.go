package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
)

func GetSystemRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/_ah/start", handlers.StartUp).Methods("GET").Name("StartUp")
	router.HandleFunc("/_ah/warmup", handlers.WarmUp).Methods("GET").Name("Warmup")

	return router
}
