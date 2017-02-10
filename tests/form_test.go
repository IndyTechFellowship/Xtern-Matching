package tests

import (
	"testing"
	"Xtern-Matching/models"
	"Xtern-Matching/handlers/services"
	"github.com/stretchr/testify/assert"
	"google.golang.org/appengine/aetest"
	"time"
)

func TestFormQuery(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if !assert.Nil(t, err, "Error instantiating context") {
		t.Fatal(err)
	}
	defer done()
	form := models.Form{"NewForm", "2017", true, []string{}}
	key, err := services.AddForm(ctx, form)
	if !assert.Nil(t, err, "Error creating form") {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	var student models.Student
	student.FirstName = "Darla"
	student.LastName = "leach"
	student.Email = "darlaleach@stockpost.com"
	student.University = "Rose-Hulman Institute of Technology"
	student.Major = "Computer Engineering"
	student.GradYear = "2017"
	student.WorkStatus = "US Citizen"
	student.HomeState = "West Virginia"
	student.Gender = "female"
	student.Skills = []models.Skill{{Name: "SQL", Category: "Database"}, {Name: "HTML", Category: "Frontend"}}
	student.Github = "https://github.com/xniccum"
	student.Linkin = ""
	student.PersonalSite = ""
	student.Interests = []string{"Product Management", "Software Engineer- Middle-tier Dev."}
	student.Grade = 5
	student.Status = "Stage 1 Approved"
	student.Resume = "public/data_mocks/sample.pdf"
	student.Active = true
	_, err = services.NewStudent(ctx, student)
	//Add a student without a form parent
	createStudent(ctx)
	time.Sleep(time.Millisecond * 500)
	if !assert.Nil(t, err, "Error adding student") {
		t.Fatal(err)
	}
	students, _, err := services.GetStudents(ctx, key)
	if !assert.Nil(t, err, "Error querying students") {
		t.Fatal(err)
	}
	if !assert.Equal(t,1, len(students), "Incorrect number of students returned on query") {
		t.Fatal()
	}
}
