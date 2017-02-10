package models

type Form struct {
	Name string 			`json:"name"`
	Year string			`json:"year"`
	Active bool			`json:"active"`
	Aliases []string		`json:"aliases"`
}
