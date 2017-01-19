package models

import "encoding/json"

// csv:"-" to exclude export
type Student struct {
	FirstName    string   `json:"firstName"`
	LastName     string   `json:"lastName"`
	Email        string   `json:"email"`
	University   string   `json:"university"`
	Major        string   `json:"major"`
	GradYear     string   `json:"gradYear"`
	WorkStatus   string   `json:"workStatus"`
	Gender       string   `json:"gender"`
	Skills       []Skill  `json:"skills"`
	Github       string   `json:"githubUrl"`
	Linkin       string   `json:"linkedinUrl"`
	PersonalSite string   `json:"personalWebiteUrl"`
	Interests    []string `json:"interests"`
	Resume       string   `json:"resume"`
	Grade        int      `json:"grade"`
	Status       string   `json:"status"`
	Active       bool     `json:"active"`
	HomeState    string   `json:"homeState"`
	//Details map[string]interface{}	`json:"details"`
	//EmailIntrest string 	`json:"interestedInEmail"`
}

/*
	Same model as Student,
	So only desired fields are encoded by JSON Marshal
 */
type StudentDecisionQuery struct {
	FirstName    string   `json:"-"`
	LastName     string   `json:"-"`
	GradYear     string   `json:"gradYear"`
	Grade        int      `json:"grade"`
	Gender       string   `json:"gender"`
	Email        string   `json:"-"`
	University   string   `json:"-"`
	Major        string   `json:"-"`
	WorkStatus   string   `json:"-"`
	Skills       []Skill  `json:"-"`
	Github       string   `json:"-"`
	Linkin       string   `json:"-"`
	PersonalSite string   `json:"-"`
	Interests    []string `json:"-"`
	Resume       string   `json:"-"`
	Status       string   `json:"-"`
	Active       bool     `json:"-"`
	HomeState    string   `json:"-"`
}

type StudentDecision StudentDecisionQuery

func (student *StudentDecision) MarshalJSON() ([]byte, error) {
	type Alias StudentDecision
	return json.Marshal(&struct {
		*Alias
		Name string	`json:"name"`
	}{(*Alias)(student),
		student.FirstName +" " + student.LastName,
	})
}

type Skill struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

func (skill Skill) MarshalCSV() ([]byte, error) {
	return []byte(skill.Name + ": " + skill.Category), nil
}
