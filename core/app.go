package main

import (
	"Xtern-Matching/routes"
	"net/http"
	"os"
	"log"
)

func init() {
	if os.Getenv("XTERN_ENVIRONMENT") != "production" {
		os.Setenv("XTERN_ENVIRONMENT","development")
		log.Println("Setting Environment to Development")
	}
	http.Handle("/", routes.NewRouter())
}
func main() {}
