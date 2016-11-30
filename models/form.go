package models

type Form struct {
	Name string 			`json:"name"`
	Aliases map[string]interface{} 	`json:"aliases"`
}
