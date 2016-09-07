package main

import (
	"encoding/json"
	"net/http"
)

func alive(res http.ResponseWriter, req *http.Request) {
	renderJSON(res, http.StatusOK, map[string]interface{}{"alive": true})
}

func sampleHandler(res http.ResponseWriter, req *http.Request) {
	renderJSON(res, http.StatusOK, map[string]interface{}{"stam": true})
}

func renderJSON(r http.ResponseWriter, status int, v interface{}) {
	var result []byte
	var err error
	result, err = json.Marshal(v)
	if err != nil {
		http.Error(r, err.Error(), 500)
		return
	}
	// json rendered fine, write out the result
	r.Header().Set("Content-Type", "application/json")
	r.WriteHeader(status)
	r.Write(result)
}
