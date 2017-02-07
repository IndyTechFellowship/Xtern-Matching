package tests

import (
	//"os"

	"Xtern-Matching/models"
	"google.golang.org/appengine/aetest"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"Xtern-Matching/handlers/services"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"encoding/json"
	"os"
	//"log"
)

func createStudent(ctx context.Context) (models.Student, error) {
	var student models.Student
	student.FirstName = "Darla"
	student.LastName = "leach"
	student.Email = "darlaleach@stockpost.com"
	student.University = "Rose-Hulman Institute of Technology"
	student.Major = "Computer Engineering"
	student.GradYear = 2017
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

	key := datastore.NewIncompleteKey(ctx, "Student", nil)
	key, err := datastore.Put(ctx, key, &student)
	if err != nil {
		return models.Student{}, err
	}
	// time.Sleep(time.Millisecond * 500)

	return student, nil
}

/*
	Creates students where each status is represented as many times as eachStatus
 */
func createStudentsWithEqualStatus(ctx context.Context, eachStatus int) ([]*datastore.Key, error) {
	statuses := [...]string{"Rejected (Stage 1)", "Rejected (Stage 2)", "Rejected (Stage 3)",
				"Undecided", "Stage 1 Approved", "Stage 2 Approved", "Stage 3 Approved"}
	var keys []*datastore.Key
	for i := 0; i < len(statuses); i++ {
		var student models.Student
		student.FirstName = "Darla"
		student.LastName = "leach"
		student.Email = "darlaleach@stockpost.com"
		student.University = "Rose-Hulman Institute of Technology"
		student.Major = "Computer Engineering"
		student.GradYear = 2017
		student.WorkStatus = "US Citizen"
		student.HomeState = "West Virginia"
		student.Gender = "female"
		student.Skills = []models.Skill{{Name: "SQL", Category: "Database"}, {Name: "HTML", Category: "Frontend"}}
		student.Github = "https://github.com/xniccum"
		student.Linkin = ""
		student.PersonalSite = ""
		student.Interests = []string{"Product Management", "Software Engineer- Middle-tier Dev."}
		student.Grade = 5
		student.Status = statuses[i]
		student.Resume = "public/data_mocks/sample.pdf"
		student.Active = true

		for j:=0; j<eachStatus; j++ {
			key := datastore.NewIncompleteKey(ctx, "Student", nil)
			key, err := datastore.Put(ctx, key, &student)
			keys = append(keys, key)
			if err != nil {
				return nil, err
			}
		}
	}
	return keys, nil
}

func TestDecisionList(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if !assert.Nil(t, err, "Error instantiating context") {
		t.Fatal(err)
	}
	defer done()
	for i:= 0; i < 10; i++ {
		createStudent(ctx)
	}
	time.Sleep(time.Millisecond * 500)
	students, err := services.GetStudentDecisionList(ctx, nil)
	if !assert.Nil(t, err, "Error getting decision list") {
		t.Fatal(err)
	}
	json.NewEncoder(os.Stdout).Encode(students)
}

func TestStatusList(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if !assert.Nil(t, err, "Error instantiating context") {
		t.Fatal(err)
	}
	defer done()
	createStudentsWithEqualStatus(ctx, 2)
	time.Sleep(time.Millisecond * 500)
	statuses := [...]string{"Rejected (Stage 1)", "Rejected (Stage 2)", "Rejected (Stage 3)",
				"Undecided", "Stage 1 Approved", "Stage 2 Approved", "Stage 3 Approved"}
	for i := 0; i < len(statuses); i++ {
		students, err := services.GetStudentsAtLeastWithStatus(ctx, statuses[i])
		if !assert.Nil(t, err, "Error getting status list") {
			t.Fatal(err)
		}
		if !assert.Equal(t, 2 * (len(statuses) - i ), len(students),
			"Invalid number of students returned in status query") {
			t.Fatal()
		}
	}
}

func TestMoveStudentsToStatus(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if !assert.Nil(t, err, "Error instantiating context") {
		t.Fatal(err)
	}
	defer done()
	keys, err := createStudentsWithEqualStatus(ctx, 2)
	if !assert.Nil(t, err, "Error creating students") {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	statuses := [...]string{"Rejected (Stage 1)", "Rejected (Stage 2)", "Rejected (Stage 3)",
				"Undecided", "Stage 1 Approved", "Stage 2 Approved", "Stage 3 Approved"}
	var testKeys []*datastore.Key
	for i := 0; i < len(statuses); i++ {
		testKeys = append(testKeys, keys[2*i])
	}
	for i := 0; i < len(statuses); i++ {
		if i+1 < len(statuses) {
			err := services.MoveStudentsToStatus(ctx, testKeys, statuses[i+1])
			if !assert.Nil(t, err, "Error Updating Student Statuses") {
				t.Fatal(err)
			}
		}
		time.Sleep(time.Millisecond * 500)
		students, err := services.GetStudentsAtLeastWithStatus(ctx, statuses[i])
		if !assert.Nil(t, err, "Error getting status list") {
			t.Fatal(err)
		}
		if !assert.Equal(t, 14 - i, len(students),
			"Invalid number of students returned in status query") {
			t.Fatal()
		}
	}
}
//
//import (
//	//"os"
//	//"Xtern-Matching/handlers/services"
//	"Xtern-Matching/models"
//	//"reflect"
//	"testing"
//	//"time"
//	"google.golang.org/appengine/aetest"
//	"github.com/stretchr/testify/assert"
//	"Xtern-Matching/handlers/services"
//	"os"
//)
//
///*
//	Creates a fresh sample student to use in testing
//*/
//func GetStudent1() models.Student {
//	var student models.Student
//	student.FirstName = "Darla"
//	student.LastName = "leach"
//	student.Email = "darlaleach@stockpost.com"
//	//student.University = "Rose-Hulman Institute of Technology"
//	//student.Major = "Computer Engineering"
//	student.GradYear = "2017"
//	//student.WorkStatus = "US Citizen"
//	//student.HomeState = "West Virginia"
//	//student.Gender = "female"
//	student.Skills = []models.Skill{{Name: "SQL", Category: "Database"}, {Name: "HTML", Category: "Frontend"}}
//	//student.Github = "https://github.com/xniccum"
//	//student.Linkin = ""
//	//student.PersonalSite = ""
//	//student.Interests = []string{"Product Management", "Software Engineer- Middle-tier Dev."}
//	//student.EmailIntrest = "false"
//	//student.R1Grade = models.Grade{Text: "C", Value: "8"}
//	student.Status = "Stage 1 Approved"
//	student.Resume = "public/data_mocks/sample.pdf"
//	student.Active = true
//	return student
//}
//
///*
//	Tests the ability to create a Student
//*/
//func TestPost(t *testing.T) {
//	ctx, done, err := aetest.NewContext()
//	if !assert.Nil(t, err, "Error instantiating context") {
//		t.Fatal(err)
//	}
//	defer done()
//
//	file, err := os.Open("resources/sample.pdf")
//	if err != nil {
//
//	}
//
//	student := GetStudent1()
//
//	// Basic input
//	_, err = services.NewStudent(ctx, student)
//	//if !assert.Nil(t, err, "Error creating student") {
//	//	t.Fatal(err)
//	//}
//	//
//	//// Check for bad input
//	//empty_student := models.Student{}
//	//status, err := services.NewStudent(ctx, empty_student)
//	//if !assert.Equal(t, status, http.StatusBadRequest, "Accepted empty student") {
//	//	t.Fatal(err)
//	//}
//}
//
///*
//	Tests both bulk retrieval and individual query.
//	Additionally ensures an invalid id can't be retrieved.
//*/
//func TestGet(t *testing.T) {
//	//ctx, done, err := aetest.NewContext()
//	//if !assert.Nil(t, err, "Error instantiating context") {
//	//	t.Fatal(err)
//	//}
//	//defer done()
//	//
//	//// Add student to get
//	//student := GetStudent1()
//	//
//	//_, err = services.NewStudent(ctx, student)
//	//if !assert.Nil(t, err, "Error creating student") {
//	//	t.Fatal(err)
//	//}
//	//
//	//student.FirstName = "Lee"
//	//student.LastName = "Robinson"
//	//student.Email = "leeson@stockpost.com"
//	//
//	//_, err = services.NewStudent(ctx, student)
//	//if !assert.Nil(t, err, "Error creating student") {
//	//	t.Fatal(err)
//	//}
//	//
//	//time.Sleep(time.Millisecond * 500)
//	//
//	//students, err := services.GetStudents(ctx)
//	//if !assert.Nil(t, err, "Error retrieving students") {
//	//	t.Fatal(err)
//	//}
//	//if !assert.Equal(t, 2, len(students), "Error in number of students retrieved") {
//	//	t.Fatal(students)
//	//}
//	//
//	//first := students[0].FirstName
//	//last := students[0].LastName
//	//student, err = services.GetStudent(ctx, id)
//	//if !assert.Nil(t, err, "Error retrieving student") {
//	//	t.Fatal(err)
//	//}
//	//if !assert.Equal(t, first, student.FirstName, "Mismatch in student first name") ||
//	//	!assert.Equal(t, last, student.LastName, "Mismatch in student last name") {
//	//	t.Fatal("Grabbed wrong student")
//	//}
//	//
//	//student, err = services.GetStudent(ctx, -4)
//	//if !reflect.DeepEqual(student, (models.Student{})) {
//	//	t.Fatal("student shouldn't exist")
//	//}
//
//}
//
//func TestResumePost(t *testing.T) {		 func TestResumePost(t *testing.T) {
//	// Due to current dependecy credentials, turning off this test until better alternative can be found		 	// Due to current dependecy credentials, turning off this test until better alternative can be found
//	//ctx, done, err := aetest.NewContext()		 	//ctx, done, err := aetest.NewContext()
//	//if err != nil {		 	//if err != nil {
//	//	t.Fatal(err)		 	//	t.Fatal(err)
//	//}		 	//}
//	//defer done()		 	//defer done()
//	//		 	//
//	//student, err := createStudent(ctx)		 	//student, err := createStudent(ctx)
//	//if err != nil {		 	//if err != nil {
//	//	t.Fatal(err)		 	//	t.Fatal(err)
//	//}		 	//}
//	//		 	//
//	//file, err := os.Open("resources/sample.pdf")		 	//file, err := os.Open("resources/sample.pdf")
//	//if err != nil {		 	//if err != nil {
//	//	t.Fatal(err)		 	//	t.Fatal(err)
//	//}		 	//}
//	//defer file.Close()		 	//defer file.Close()
//	//		 	//
//	//err = services.UpdateResume(ctx, student.Id, file)		 	//err = services.UpdateResume(ctx, student.Id, file)
//	//if err != nil {		 	//if err != nil {
//	//	t.Fatal(err)		 	//	t.Fatal(err)
//	//}		 	//}
//}
