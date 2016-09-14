package models

import "gopkg.in/mgo.v2/bson"

type Company struct {
	Id int 			`json:"id"`
	Name string 		`json:"name"`
	Password string  	`json:"password"`
	Positions Postings 	`json:"postings"`
}

type Companies []Company