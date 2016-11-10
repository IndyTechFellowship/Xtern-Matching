package models

type Student struct {
	FirstName string    		`json:"firstName"`
	LastName string     		`json:"lastName"`
	Email string	    		`json:"email"`
	Details map[string]interface{}
	//University string   `json:"university"`
	//Major string 	    `json:"major"`
	GradYear string     		`json:"gradYear"`
	//WorkStatus string   `json:"workStatus"`
	//HomeState string    `json:"homeState"`
	//Gender string       `json:"gender"`
	Skills []Skill	    		`json:"languages"`
	//Github string       `json:"githubUrl"`
	//Linkin string       `json:"linkedinUrl"`
	//PersonalSite string `json:"personalWebiteUrl"`
	//Interests []string  `json:"interestedIn"`
	//EmailIntrest string 	`json:"interestedInEmail"`
	Grade map[string]int		`json:"r1Grade"`
	Status string	    		`json:"status"`
	Active bool	    		`json:"active"`
	Resume string	    		`json:"resume"`
}

type Skill struct {
	Name string 		`json:"name"`
	Category string		`json:"category"`
}