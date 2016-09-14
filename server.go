package main

import (
	"net/http"
	"html/template"
	//"log" log.Print()
)

func init() {
	//main()
	http.HandleFunc("/",func(w http.ResponseWriter, req *http.Request) {
		t, _ := template.ParseFiles("public/index.html")  // Parse template file.
		t.Execute(w, nil)  // merge.
	})
	
	//Ignore request for now
	http.HandleFunc("/favicon.ico",func(w http.ResponseWriter, req *http.Request) {
		//t, _ := template.ParseFiles("public/index.html")  // Parse template file.
		//t.Execute(w, nil)  // merge.
	})
	
	//TODO: resolve lib folder
	http.HandleFunc("/semantic/",func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "public"+req.URL.Path)
	})
	
	http.HandleFunc("/css/",func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "public"+req.URL.Path)
		
	})
	
	http.HandleFunc("/js/",func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "public"+req.URL.Path)
	})
	
	http.HandleFunc("/node_modules/",func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "public"+req.URL.Path)
	})
}
