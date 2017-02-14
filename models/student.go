package models


import "encoding/json"
import (
	"google.golang.org/appengine/datastore"
)


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

	Grade        float64  `json:"grade"`
	ReviewerGrades	[]ReviewerGrade	`json:"reviewerGrades"`

	Status       string   `json:"status"`
	Active       bool     `json:"active"`
	HomeState    string   `json:"homeState"`
	Ethnicity    string   `json:"ethnicity"`
	//Details map[string]interface{}	`json:"details"`
	//EmailIntrest string 	`json:"interestedInEmail"`
}

/*
	Same model as Student,
	So only desired fields are encoded by JSON Marshal
 */
type StudentDecisionQuery struct {
	Id int64              `json:"key" datastore:"-"`
	FirstName    string   `json:"-"`
	LastName     string   `json:"-"`
	GradYear     string   `json:"gradYear"`
	Grade        float64  `json:"grade"`
	Gender       string   `json:"gender"`
	WorkStatus   string   `json:"workStatus"`
	Ethnicity    string   `json:"ethnicity"`
}

type StudentDecision StudentDecisionQuery

func (student *StudentDecision) MarshalJSON() ([]byte, error) {
	type Alias StudentDecision
	return json.Marshal(&struct {
		*Alias
		Name string	`json:"name"`
	}{(*Alias)(student),
		student.FirstName + " " + student.LastName,
	})
}

type Skill struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

type ReviewerGrade struct {
	Reviewer	*datastore.Key `json:"reviewer"`
	Grade 		int `json:"grade"`
}

func (skill Skill) MarshalCSV() ([]byte, error) {
	return []byte(skill.Name + ": " + skill.Category), nil
}
