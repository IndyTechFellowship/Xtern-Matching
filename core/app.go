package main

import (
	"Xtern-Matching/routes"
	"net/http"
	"os"
)

func init() {
	if os.Getenv("XTERN_ENVIRONMENT") != "production" {
		os.Setenv("XTERN_ENVIRONMENT","development")
	}

	http.Handle("/", routes.NewRouter())
}
func main() {}
