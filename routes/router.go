package routes

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router = SetStudentRoutes(router)
	router = SetAuthenticationRoutes(router)
	router = SetAdminRoutes(router)
	return router
}
