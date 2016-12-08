package models

import (
	"google.golang.org/appengine/datastore"
	"github.com/pkg/errors"
)

type Organization struct {
	Name string 		  `json:"name"`
	Kind string		  `json:"kind"`
	Students []*datastore.Key `json:"students"`
}

func NewOrganization(name string, kind string) Organization {
	students := make([]*datastore.Key,0)
	return Organization{Name: name, Kind: kind, Students: students}
}

func (org *Organization) AddStudent(studentKey *datastore.Key) bool {
	for _, key := range org.Students {
		if key.Equal(studentKey) {
			return false
		}
	}
	org.Students = append(org.Students,studentKey)
	return true
}

func (org *Organization) RemoveStudent(studentKey *datastore.Key) bool {
	for i, key := range org.Students {
		if key.Equal(studentKey) {
			org.Students = append(org.Students[:i], org.Students[i+1:]...)
			return true
		}
	}
	return false
}

func (org *Organization) MoveStudent(studentKey *datastore.Key, pos int) error {
	if org.RemoveStudent(studentKey) {
		for i := range org.Students {
			if pos == i {
				leftStudents := append(org.Students[:i], studentKey)
				org.Students = append(leftStudents, org.Students[i:]...)
				return nil
			}
		}
	}
	return errors.New("Organization does not currently have student in list")
}

