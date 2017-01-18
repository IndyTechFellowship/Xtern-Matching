package models

type ReviewGroup struct {
	Reviewers       []*User  `json:"reviewers"`
	Students []*Student 	`json:"students"`
}
