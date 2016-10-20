package main

import (
	"net/http"
	"Xtern-Matching/routes"
	"os"
)

func init() {
	if os.Getenv("XTERN_ENVIRONMENT") != "production" {
		os.Setenv("XTERN_ENVIRONMENT","development")
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS","environments/development/cloudstore-dev.json")
	}

	http.Handle("/", routes.NewRouter())
}
func main() {}