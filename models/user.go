package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	emailID  string  `json:"email" bson:"email"`
	password string  `json:"password" bson:"password"`
}

type Users []User
