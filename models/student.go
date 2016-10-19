package models

type Student struct {
	Id int64            `json:"_id" datastore:"-"`
	FirstName string    `json:"firstName"`
	LastName string     `json:"lastName"`
	Email string	    `json:"email"`
	University string   `json:"university"`
	Major string 	    `json:"major"`
	GradYear string     `json:"gradYear"`
	WorkStatus string   `json:"workStatus"`
	HomeState string    `json:"homeState"`
	Gender string       `json:"gender"`
	Skills []Skill	    `json:"languages"`
	Github string       `json:"githubUrl"`
	Linkin string       `json:"linkedinUrl"`
	PersonalSite string `json:"personalWebiteUrl"`
	Interests []string  `json:"interestedIn"`
	EmailIntrest string `json:"interestedInEmail"`
	R1Grade Grade	    `json:"r1Grade"`
	Status string	    `json:"status"`
	Active bool			`json:"active"`
	Resume string		`json:"resume"`
	Comments []Comment  `json:"comments"`
}

type Skill struct {
	Name string 		`json:"name"`
	Category string		`json:"category"`
}

type Grade struct {
	Text string 		`json:"text"`
	Value string		`json:"value"`
}

type Comment struct {
	Author string 		`json:"author"`
	Group string		`json:"group"`
	Text string 		`json:"Text"`
}