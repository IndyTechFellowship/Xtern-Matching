package main

import (
	"net/http"
	"Xtern-Matching/routes"
)

func init() {

	http.Handle("/", routes.NewRouter())
}
func main() {}