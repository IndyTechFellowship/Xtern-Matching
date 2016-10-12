package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
)

func GetAdminRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/admin/register", handlers.Register).Methods("POST").Name("Register")
	router.HandleFunc("/admin/bulkregister", handlers.BulkRegister).Methods("POST").Name("BulkRegister")
	router.HandleFunc("/admin/getusers/{Role}/{Organization}", handlers.GetUsers).Methods("GET").Name("GetUsers")
	// Fix this later - Should look like this
	//router.HandleFunc("/admin/getusers/?Role={Role}&Organization={Organization}", handlers.GetUsers).Methods("GET").Name("GetUsers")
	return router
}