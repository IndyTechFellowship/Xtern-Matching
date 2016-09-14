package handlers

import (
	"net/http"
	"html/template"
)

func Index(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("public/index.html")
	t.Execute(w, nil)
}
