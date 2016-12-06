package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
)

func GetAdminRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/admin/register", handlers.Register).Methods("POST").Name("Register")
	router.HandleFunc("/admin/bulkregister", handlers.BulkRegister).Methods("POST").Name("BulkRegister")
	router.HandleFunc("/admin/getUser", handlers.GetUser).Methods("GET").Name("GetUser")
	router.HandleFunc("/admin/getusers/{Role}/{Organization}", handlers.GetUsers).Methods("GET").Name("GetUsers")
	router.HandleFunc("/admin", handlers.PutUser).Methods("PUT").Name("UpdateUser")
	router.HandleFunc("/admin/{Id}", handlers.DeleteUser).Methods("DELETE").Name("DeleteUser")
	return router
}