package models

import "gopkg.in/mgo.v2/bson"

type Company struct {
	Id bson.ObjectId 	`json:"id" bson:"_id,omitempty"`
	Name string 		`json:"name" bson:"name"`
	Password string  	`json:"password" bson:"password"`
	Positions Postings 	`json:"postings" bson:"postings,omitempty"`
}

type Companies []Company