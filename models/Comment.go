package models

import "google.golang.org/appengine/datastore"

type Comment struct {
	Author datastore.Key	`json:"author"`
	Text string 		`json:"Text"`
}
