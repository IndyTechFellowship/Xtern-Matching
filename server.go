package main

import (
	"net/http"
	"html/template"
)

func init() {
	//main()
	http.HandleFunc("/",func(w http.ResponseWriter, req *http.Request) {
		t, _ := template.ParseFiles("public/index.html")  // Parse template file.
		t.Execute(w, nil)  // merge.
	})
}
