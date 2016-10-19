package routes

import (
	"github.com/gorilla/mux"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/codegangsta/negroni"
)

func NewRouter() *mux.Router {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//return context.Get(r,"ctx")
			return []byte("My Secret"), nil
		},
		//SigningMethod: jwt.GetSigningMethod("AppEngine"),
		// SigningMethod: jwt.SigningMethodHS512,
	});
	router := mux.NewRouter().StrictSlash(true)

	router.PathPrefix("/auth").Handler(negroni.New(
		negroni.Wrap(GetAuthenticationRoutes(mux.NewRouter().StrictSlash(true))),
	))
	router.PathPrefix("/admin").Handler(negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(GetAdminRoutes(mux.NewRouter().StrictSlash(true))),
	))
	router.PathPrefix("/student").Handler(negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(GetStudentRoutes(mux.NewRouter().StrictSlash(true))),
	))
	router.PathPrefix("/company").Handler(negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(GetCompanyRoutes()),
	))

	return router
}
