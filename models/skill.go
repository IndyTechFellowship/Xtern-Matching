package models

import "gopkg.in/mgo.v2/bson"

type Skill struct {
	Id bson.ObjectId 	`json:"id"`
	Name string 		`json:"name"`
	Category string		`json:"category"`
}

type Skills []Skill
