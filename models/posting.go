package models

import "gopkg.in/mgo.v2/bson"

type Posting struct {
	Id bson.ObjectId	`json:"id" bson:"_id,omitempty"`
	Title string 		`json:"title" bson:"title"`
	Description string	`json:"description" bson:"description"`
	skills []string		`json:"skils" bson:"skills"`
}

type Postings []Posting
