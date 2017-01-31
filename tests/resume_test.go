package tests

import (
	"testing"
	"google.golang.org/appengine/aetest"
	"github.com/stretchr/testify/assert"
	"Xtern-Matching/models"
	"Xtern-Matching/handlers/services"
	"os"
	"time"
)

func TestPost(t *testing.T) {
	os.Setenv("XTERN_ENVIRONMENT", "")
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS",
		"../core/environments/development/cloudstore-dev.json")
	ctx, done, err := aetest.NewContext()
	if !assert.Nil(t, err, "Error instantiating context") {
		t.Fatal(err.Error())
	}
	defer done()

	// Ideally would have a helper method for this, currently sitting in another branch
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
	student.Active = true

	// NewStudent should return the key to enable a reference to it to enable proper testing
	_, err = services.NewStudent(ctx, student)
	if !assert.Nil(t, err, "Error adding student") {
		t.Fatal(err.Error())
	}
	time.Sleep(time.Millisecond * 500)
	buf, err := services.ExportAllResumes(ctx)
	if !assert.Nil(t, err, "Error exporting student resumes") {
		t.Fatal(err.Error())
	}
	file, err := os.Create("manual/archive.zip")
	if !assert.Nil(t, err, "Error creating new file") {
		t.Fatal(err.Error())
	}
	defer file.Close()

	_, err = file.Write(buf.Bytes())
	if !assert.Nil(t, err, "Error writing to file") {
		t.Fatal(err.Error())
	}
}