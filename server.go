package main

import (
	"net/http"
	"routes"
)

func init() {
	http.Handle("/", routes.NewRouter())
}