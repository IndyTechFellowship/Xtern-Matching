package models

import "google.golang.org/appengine/datastore"

type Comment struct {
	Message string		`json:"message"`
	Author datastore.Key	`json:"Author"`
}

