package models

type User struct {
	Email string 		`json:"email"`
	Password string		`json:"password"`
	Organization string	`json:"organization"`
	Role string		`json:"role"`
}
