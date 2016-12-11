package models

import "google.golang.org/appengine/datastore"

type Comment struct {
	Message string		`json:"message"`
	AuthorName string 	`json:"authorName"`
	Author *datastore.Key	`json:"author"`
}

