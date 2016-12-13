package models

type Student struct {
	Id           int64     `json:"_id" datastore:"-" csv:"-"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Email        string    `json:"email"`
	University   string    `json:"university"`
	Major        string    `json:"major"`
	GradYear     string    `json:"gradYear"`
	WorkStatus   string    `json:"workStatus"`
	HomeState    string    `json:"homeState"`
	Gender       string    `json:"gender"`
	Skills       []Skill   `json:"languages"`
	Github       string    `json:"githubUrl"`
	Linkin       string    `json:"linkedinUrl"`
	PersonalSite string    `json:"personalWebiteUrl"`
	Interests    []string  `json:"interestedIn"`
	EmailIntrest string    `json:"interestedInEmail"`
	R1Grade      Grade     `json:"r1Grade"`
	Status       string    `json:"status"`
	Active       bool      `json:"active"`
	Resume       string    `json:"resume"`
	Comments     []Comment `json:"comments" csv:"-"`
}

type Skill struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

type Grade struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

func (skill Skill) MarshalCSV() ([]byte, error) {
	return []byte(skill.Name + ": " + skill.Category), nil
}

func (grade Grade) MarshalCSV() ([]byte, error) {
	return []byte(grade.Value + ": " + grade.Text), nil
}

type Comment struct {
	Author string `json:"author"`
	Group  string `json:"group"`
	Text   string `json:"Text"`
}
