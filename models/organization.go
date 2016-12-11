package models

import (
	"google.golang.org/appengine/datastore"
	"log"
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

func (org *Organization) ContainStudent(studentKey *datastore.Key) int {
	for i, key := range org.Students {
		if key.Equal(studentKey) {
			return i
		}
	}
	return -1
}

func (org *Organization) RemoveStudent(studentKey *datastore.Key) {
	for i, key := range org.Students {
		if key.Equal(studentKey) {
			org.Students = append(org.Students[:i], org.Students[i+1:]...)
			return
		}
	}
}

func (org *Organization) MoveStudent(studentKey *datastore.Key, pos int) {
	students := make([]*datastore.Key,len(org.Students))
	i := 0
	j := 0
	for i < len(org.Students) && j < len(students) {
		if pos == j {
			log.Printf("Posisiton=%v", i)
			students[j] = studentKey
			j++
		} else if org.Students[i].Equal(studentKey) {
			log.Printf("Remove=%v", i)
			i++
		} else {
			students[j] = org.Students[i]
			i++
			j++
		}
	}
	org.Students = students
}

