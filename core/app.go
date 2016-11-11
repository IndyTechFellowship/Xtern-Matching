package main

import (
	"net/http"
	"Xtern-Matching/routes"
	"os"
)

func init() {
	if os.Getenv("XTERN_ENVIRONMENT") != "production" {
		os.Setenv("XTERN_ENVIRONMENT","development")
	}

	http.Handle("/", routes.NewRouter())
}
func main() {}