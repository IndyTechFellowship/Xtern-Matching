package models

import (
	"google.golang.org/appengine/datastore"
)

type ReviewGroup struct {
	Reviewers       []*datastore.Key  `json:"reviewers"`
	Students []*datastore.Key 	`json:"students"`
}
