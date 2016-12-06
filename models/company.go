package models

type Company struct {
	Id int64            `json:"_id" datastore:"-"`
	Name string    		`json:"name"`
	StudentIds []int64	`json:"studentIds"`
}