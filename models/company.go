package models

type Company struct {
	Name string    		`json:"name"`
	Type string		`json:"type"`
	StudentIds []int64	`json:"studentIds"`
}