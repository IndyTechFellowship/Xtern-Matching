package models

import "gopkg.in/mgo.v2/bson"

type Skill struct {
	Id bson.ObjectId 	`json:"id" bson:"_id,omitempty"`
	Name string 		`json:"name" bson:"name"`
	Tag string		`json:"tag" bson:"tag"`
}

type Skills []Skill
