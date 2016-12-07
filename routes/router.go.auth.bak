package routes

/**
	DO NOT COMMIT THIS FILE AS A GO FILE
	
	USE THIS FILE TO BYPASS AUTH AND RUN SCRIPTS
	-- will mostlikely cause front end errors

	TO USE:
	rename router.go to router.go.primary.bak
	rename this file to router.go

	WHEN FINISHED:
	rename this file to router.go.bak
	remane router.go.primary.bak to router.go 

**/

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router = GetAdminRoutes(router) // Add Admin Routes
	router = GetStudentRoutes(router) // Add Student Routes
	router = GetAuthenticationRoutes(router) // Add Authentication Routes
	return router
}
