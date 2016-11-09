package main

import (
	"Xtern-Matching/routes"
	"net/http"

	"os"

	"google.golang.org/appengine"
)

func init() {
	if os.Getenv("XTERN_ENVIRONMENT") != "production" {
		os.Setenv("XTERN_ENVIRONMENT", "development")
		// os.Setenv("GOOGLE_APPLICATION_CREDENTIALS","environments/development/cloudstore-dev.json")
	}

	//Seed Datasbase

	http.Handle("/", routes.NewRouter())
}
func main() {
	appengine.Main()
}
