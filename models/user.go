package models

type User struct {
	Email string 		`json:"email"`
	Password []byte		`json:"password"`
	Organization string	`json:"organization"`
	Role string		`json:"role"`
}
