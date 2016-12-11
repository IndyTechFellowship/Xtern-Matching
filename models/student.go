package models

type Student struct {
	FirstName string    		`json:"firstName"`
	LastName string     		`json:"lastName"`
	Email string	    		`json:"email"`
	University string   		`json:"university"`
	Major string 	    		`json:"major"`
	GradYear string     		`json:"gradYear"`
	WorkStatus string   		`json:"workStatus"`
	Gender string       		`json:"gender"`
	Skills []Skill	    		`json:"skills"`
	Github string       		`json:"githubUrl"`
	Linkin string       		`json:"linkedinUrl"`
	PersonalSite string 		`json:"personalWebiteUrl"`
	Interests []string  		`json:"interests"`
	Resume string	    		`json:"resume"`
	Grade int			`json:"grade"`
	Status string	    		`json:"status"`
	Active bool	    		`json:"active"`
	HomeState string    `json:"homeState"`
	//Details map[string]interface{}	`json:"details"`
	//EmailIntrest string 	`json:"interestedInEmail"`
}

type Skill struct {
	Name string 		`json:"name"`
	Category string		`json:"category"`
}