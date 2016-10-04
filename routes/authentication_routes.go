package routes

import (
	"github.com/gorilla/mux"
	"Xtern-Matching/handlers"
	_ "github.com/someone1/gcp-jwt-go"
	//"net/http"
	//"google.golang.org/appengine"
	//"github.com/gorilla/context"
)

//func AppEngineContextGrab(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
//	ctx := appengine.NewContext(r)
//	context.Set(r,"ctx",ctx)
//	next(rw, r)
//	context.Clear(r)
//}

func GetAuthenticationRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/auth/login", handlers.Login).Methods("POST").Name("Login")
	router.HandleFunc("/auth/logout", handlers.Logout).Methods("POST").Name("Logout")
	return router
}
