package models

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"google.golang.org/appengine/datastore"
	"github.com/pkg/errors"
)

type Organization struct {
	Name string 		`json:"name"`
	Type string		`json:"type"`
	Students arraylist.List `json:"students"`
}

func (org *Organization) AddStudent(studentKey datastore.Key) bool {
	if !org.Students.Contains(studentKey) {
		org.Students.Add(studentKey)
		return true
	}
	return false
}

func (org *Organization) RemoveStudent(studentKey datastore.Key) {
	org.Students.Remove(studentKey)
}

func (org *Organization) MoveStudent(studentKey datastore.Key, pos int) error {
	if !org.Students.Contains(studentKey) {
		return errors.New("Organization does not currently have student in list")
	}
	position, _ := org.Students.Find(func(index int, key datastore.Key) {
		return key == studentKey
	})
	org.Students.Remove(position)
	org.Students.Insert(pos,studentKey)
	return nil
}

