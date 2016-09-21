package models

type Student struct {
	FirstName string    `json:"firstName"`
	LastName string     `json:"lastName"`
	Email string	    `json:"email"`
	University string   `json:"university"`
	Major string 	    `json:"major"`
	GradYear string     `json:"gradYear"`
	WorkStatus string   `json:"workStatus"`
	HomeState string    `json:"homeState"`
	Gender string       `json:"gender"`
	Skills []string	    `json:"skills"`
	Github string       `json:"githubUrl"`
	Linkin string       `json:"linkedinUrl"`
	PersonalSite string `json:"personalWebiteUrl"`
	Interests []string  `json:"interestedIn"`
	EmailIntrest string `json:"interestedInEmail"`
	Status string	    `json:"status"`
}