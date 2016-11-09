package main

import (
	"net/http"
	"Xtern-Matching/routes"

	"os"
	"google.golang.org/appengine"
)

func init() {
	if os.Getenv("XTERN_ENVIRONMENT") != "production" {
		os.Setenv("XTERN_ENVIRONMENT","development")
		// os.Setenv("GOOGLE_APPLICATION_CREDENTIALS","environments/development/cloudstore-dev.json")
	}

	//Seed Datasbase

	http.Handle("/", routes.NewRouter())
}
func main() {
	appengine.Main()
}