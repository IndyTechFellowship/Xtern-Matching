package models

import (
	"google.golang.org/appengine/datastore"
	"log"
)

type Organization struct {
	Name string 		  `json:"name"`
	StudentRanks []StudentRank `json:"students"`
}

type StudentRank struct {
	Student	*datastore.Key `json:"student"`
	Rank 		int `json:"rank"`
}

func NewOrganization(name string) Organization {
	studentRanks := make([]StudentRank,0)
	return Organization{Name: name, StudentRanks: studentRanks}
}

func (org *Organization) AddStudent(studentKey *datastore.Key) bool {
	for _, studentRank := range org.StudentRanks {
		if studentRank.Student.Equal(studentKey) {
			return false
		}
	}
	var newStudentRank StudentRank
		newStudentRank.Student = studentKey
		newStudentRank.Rank = len(org.StudentRanks)
	org.StudentRanks = append(org.StudentRanks,newStudentRank)
	return true
}

func (org *Organization) ContainStudent(studentKey *datastore.Key) int {
	for i, studentRank := range org.StudentRanks {
		if studentRank.Student.Equal(studentKey) {
			return i
		}
	}
	return -1
}

func (org *Organization) RemoveStudent(studentKey *datastore.Key) {
	for i, studentRank := range org.StudentRanks {
		if studentRank.Student.Equal(studentKey) {
			org.StudentRanks = append(org.StudentRanks[:i], org.StudentRanks[i+1:]...)
			return
		}
	}
}

func (org *Organization) MoveStudent(studentKey *datastore.Key, pos int) {
	studentRanks := make([]StudentRank,len(org.StudentRanks))
	i := 0
	j := 0
	for i < len(org.StudentRanks) && j < len(studentRanks) {
		if pos == j {
			log.Printf("Posisiton=%v", i)
			studentRanks[j].Student = studentKey
			j++
		} else if org.StudentRanks[i].Student.Equal(studentKey) {
			log.Printf("Remove=%v", i)
			i++
		} else {
			studentRanks[j] = org.StudentRanks[i]
			i++
			j++
		}
	}
	org.StudentRanks = studentRanks
}

