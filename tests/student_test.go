package tests

import (
	"os"
	"strings"
	"testing"
	"encoding/json"
	"google.golang.org/appengine/aetest"
	"Xtern-Matching/handlers/services"
	"Xtern-Matching/models"
	"golang.org/x/net/context"
)

func TestPost(t *testing.T) {
	//ctx, done, err := aetest.NewContext()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//defer done()
}

func TestGet(t *testing.T) {
	//ctx, done, err := aetest.NewContext()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//defer done()
}

/*
	Creates a sample user in the test database we can use for purposes such as testing a resume upload
*/
func createUser(ctx context.Context) (models.Student, error) {
	userJson := "{\"_id\":123456789012345,\"firstName\":\"Verna\",\"lastName\":\"Gomez\",\"email\":\"vernagomez@sportan.com\",\"university\":\"Rose-Hulman Institute of Technology\",\"major\":\"Computer Science\",\"gradYear\":\"2019\",\"workStatus\":\"EAD\",\"homeState\":\"Minnesota\",\"gender\":\"female\",\"languages\":[{\"name\":\"Ruby\",\"category\":\"General\"},{\"name\":\"Python\",\"category\":\"General\"},{\"name\":\"HTML\",\"category\":\"Front-End\"},{\"name\":\"CSS\",\"category\":\"Front-End\"},{\"name\":\"Knockout\",\"category\":\"Front-End\"},{\"name\":\"Ruby\",\"category\":\"General\"},{\"name\":\".Net\",\"category\":\"Full-Stack\"},{\"name\":\"SQL\",\"category\":\"Database\"}],\"githubUrl\":\"https://github.com/davisnygren\",\"linkedinUrl\":\"http://www.linkedin.com/\",\"personalWebiteUrl\":\"http://www.rose-hulman.edu/\",\"interestedIn\":[\"Product Management\",\"Project Management\",\"Software Engineer- Front-end Web Dev\"],\"interestedInEmail\":\"true\",\"r1Grade\":{},\"status\":\"Stage 1 Approved\",\"comments\":[{\"author\":\"Adams Lane\",\"group\":\"Xtern\",\"Text\":\"Tempor Lorem elit nostrud pariatur in non quis. Laboris minim incididunt in ad mollit anim cupidatat enim commodo proident nostrud minim excepteur. Deserunt duis mollit amet sint aliquip nulla ea aliquip labore tempor mollit nulla. Nostrud anim deserunt anim ex aute. Fugiat officia irure do excepteur occaecat exercitation. Sunt sit enim reprehenderit nostrud minim.\"}]}"
	var student models.Student
	decoder := json.NewDecoder(strings.NewReader(userJson))
	if err := decoder.Decode(&student); err != nil {
		return student, err
	}
	_,err := services.NewStudent(ctx,&student)
	if err != nil {
		return student, err
	}
	return student, nil
}

func TestResumePost(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()
	
	student, err := createUser(ctx)
	if err != nil {
		t.Fatal(err)
	}
	
	file, err := os.Open("resources/sample.pdf")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	
	err = services.UpdateResume(ctx, student.Id, file)
	if err != nil {
		t.Fatal(err)
	}	
}