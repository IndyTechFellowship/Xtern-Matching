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

func GetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/auth/login", handlers.Login).Methods("POST").Name("Login")
	return router
}
