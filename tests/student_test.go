package tests

import (
	//"os"
	"testing"
	"google.golang.org/appengine/aetest"
	"Xtern-Matching/handlers/services"
	"net/http"
	"Xtern-Matching/models"
	"google.golang.org/appengine/datastore"
	"time"
	"reflect"
	"golang.org/x/net/context"
	"github.com/stretchr/testify/assert"
)


func createStudent(ctx context.Context) (models.Student, error) {
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
	student.Skills = []models.Skill{ {Name: "SQL",Category:"Database"}, {Name: "HTML",Category:"Frontend"}}
	student.Github = "https://github.com/xniccum"
	student.Linkin = ""
	student.PersonalSite = ""
	student.Interests = []string{"Product Management","Software Engineer- Middle-tier Dev."}
	student.EmailIntrest = "false"
	student.R1Grade = models.Grade{Text: "C",Value: "8"}
	student.Status = "Stage 1 Approved"
	student.Comments = []models.Comment{}
	student.Resume = "public/data_mocks/sample.pdf"
	student.Active = true

	key := datastore.NewIncompleteKey(ctx, "Student", nil)
	key, err := datastore.Put(ctx, key, &student);
	if err != nil {
		return models.Student{}, err
	}
	time.Sleep(time.Millisecond * 500)
	student.Id = key.IntID()

	return student, nil
}

func TestPost(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

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
	student.Skills = []models.Skill{ {Name: "SQL",Category:"Database"}, {Name: "HTML",Category:"Frontend"}}
	student.Github = "https://github.com/xniccum"
	student.Linkin = ""
	student.PersonalSite = ""
	student.Interests = []string{"Product Management","Software Engineer- Middle-tier Dev."}
	student.EmailIntrest = "false"
	student.R1Grade = models.Grade{ Text: "C",Value: "8"}
	student.Status = "Stage 1 Approved"
	student.Comments = []models.Comment{}


	// Basic input
	_, err = services.NewStudent(ctx, student)
	if err != nil {
		t.Fatal(err)
	}

	// Check for bad input
	empty_student := models.Student{}
	status, err := services.NewStudent(ctx, empty_student)
	if status != http.StatusBadRequest {
		t.Fatal("Accepted empty student")
	}
}

func TestGet(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	// Add student to get
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
	student.Skills = []models.Skill{ {Name: "SQL",Category:"Database"}, {Name: "HTML",Category:"Frontend"}}
	student.Github = "https://github.com/xniccum"
	student.Linkin = ""
	student.PersonalSite = ""
	student.Interests = []string{"Product Management","Software Engineer- Middle-tier Dev."}
	student.EmailIntrest = "false"
	student.R1Grade = models.Grade{Text: "C",Value: "8"}
	student.Status = "Stage 1 Approved"
	student.Comments = []models.Comment{}

	key := datastore.NewIncompleteKey(ctx, "Student", nil)
	if _, err := datastore.Put(ctx, key, &student); err != nil {
		t.Fatal(err)
	}


	student.FirstName = "Lee"
	student.LastName = "Robinson"
	student.Email = "leeson@stockpost.com"

	key = datastore.NewIncompleteKey(ctx, "Student", nil)
	if _, err := datastore.Put(ctx, key, &student); err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Millisecond * 500)

	students, err := services.GetStudents(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(students) != 2 {
		//log.Print(len(students))
		//t.Fatal("Didn't grab all students")
		t.Fatal(students)
	}

	first := students[0].FirstName
	last := students[0].LastName
	id := students[0].Id
	student, err = services.GetStudent(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
	if student.FirstName != first || student.LastName != last {
		t.Fatal("Grabbed wrong student")
	}

	student, err = services.GetStudent(ctx, -4)
	if !reflect.DeepEqual(student, (models.Student{})) {
		t.Fatal("student shouldn't exist")
	}

}

func TestResumePost(t *testing.T) {
	// Due to current dependecy credentials, turning off this test until better alternative can be found
	//ctx, done, err := aetest.NewContext()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//defer done()
	//
	//student, err := createStudent(ctx)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//file, err := os.Open("resources/sample.pdf")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//defer file.Close()
	//
	//err = services.UpdateResume(ctx, student.Id, file)
	//if err != nil {
	//	t.Fatal(err)
	//}
}
