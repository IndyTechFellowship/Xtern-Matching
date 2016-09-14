package main

import (
	"net/http"
	"html/template"
	//"appengine"
	//"appengine/user"
	"Xtern-Matching/routes"
)

func init() {
	http.Handle("/",routes.NewRouter())

	http.HandleFunc("/",func(w http.ResponseWriter, req *http.Request) {
		t, _ := template.ParseFiles("public/index.html")
		t.Execute(w, nil)
	})
}
