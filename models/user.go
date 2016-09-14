package models

type Student struct {
	ID  int		    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName string     `json:"lastName"`
	Email string	    `json:"email"`
	University string   `json:"university"`
	Major string 	    `json:"major"`
	GradYear string     `json:"gradYear"`
	WorkStatus string   `json:"workStatus"`
	HomeState string    `json:"homeState"`
	Gender string       `json:"gender"`
	Skills Skills	    `json:"skills"`
	Github string       `json:"githubUrl"`
	Linkin string       `json:"linkedinUrl"`
	PersonalSite string `json:"personalWebiteUrl"`
	Interests []string  `json:"interestedIn"`
	EmailIntrest bool   `json:"interestedInEmail"`
	Status bool	    `json:"status"`
}
