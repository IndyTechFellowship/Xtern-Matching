package models

import "gopkg.in/mgo.v2/bson"

type Posting struct {
	Id bson.ObjectId	`json:"id"`
	Title string 		`json:"title"`
	Description string	`json:"description"`
	skills []string		`json:"skills"`
}

type Postings []Posting
