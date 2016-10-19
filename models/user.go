package models

type User struct {
	Id int64            `json:"_id" datastore:"-"`
	Name string 		`json:"name"`
	Email string 		`json:"email"`
	Password string		`json:"password"`
	Organization string	`json:"organization"`
	Role string		`json:"role"`
}
