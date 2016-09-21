package models

type Comment struct {
	Author string 		`json:"author"`
	Group string		`json:"group"`
	Text string 		`JSON:"Text"`
}
