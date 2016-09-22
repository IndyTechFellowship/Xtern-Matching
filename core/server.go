package main

import (
	"Xtern-Matching/routes"
	"net/http"
)

func init() {
	http.Handle("/", routes.NewRouter())
}
func main() {

}