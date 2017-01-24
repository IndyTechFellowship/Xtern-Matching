package routes

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options {
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("My Secret"), nil
		},
		SigningMethod: jwt.SigningMethodHS512,
	})
	router := mux.NewRouter().StrictSlash(true)

	router.PathPrefix("/_ah").Handler(negroni.New(
		negroni.Wrap(GetSystemRoutes(mux.NewRouter().StrictSlash(true))),
	))

	router.PathPrefix("/auth").Handler(negroni.New(
		negroni.Wrap(GetAuthenticationRoutes(mux.NewRouter().StrictSlash(true))),
	))
	router.PathPrefix("/comment").Handler(negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(GetCommentRoutes(mux.NewRouter().StrictSlash(true))),
	))
	router.PathPrefix("/organization").Handler(negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(GetOrganizationRoutes(mux.NewRouter().StrictSlash(true))),
	))
	router.PathPrefix("/student").Handler(negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(GetStudentRoutes(mux.NewRouter().StrictSlash(true))),
	))
	router.PathPrefix("/user").Handler(negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(GetUserRoutes(mux.NewRouter().StrictSlash(true))),
	))

	return router
}
