package models

type User struct {
	Email string 		`json:"email"`
	Password string		`json:"password"`
	Organization int64	`json:"organization"`
	Role string		`json:"role"`
}
