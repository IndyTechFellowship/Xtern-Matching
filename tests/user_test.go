package tests

import (
	"testing"
	"google.golang.org/appengine/aetest"
	"Xtern-Matching/handlers/services"
	"Xtern-Matching/models"
	"net/http"
)

func TestRegister(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	user1 := models.User{}
	user1.Email = "xniccum@gmail.com"
	user1.Password = "admin1"
	user1.Organization = "Xtern"
	user1.Role = "admin"

	user2 := models.User{}
	user2.Email = "samael@work.com"
	user2.Password = "work"
	user2.Organization = "Xtern"
	user2.Role = "company"

	responseStatus, err := services.Register(ctx,user1)
	if responseStatus != http.StatusAccepted {
		t.Fail()
	}
	responseStatus, err = services.Register(ctx,user2)
	if responseStatus != http.StatusAccepted {
		t.Fail()
	}

	responseStatus, err = services.Register(ctx,user1)
	if responseStatus != http.StatusInternalServerError {
		t.Fail()
	}
}

func TestLogin(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	user1 := models.User{}
	user1.Email = "xniccum@gmail.com"
	user1.Password = "admin1"
	user1.Organization = "Xtern"
	user1.Role = "admin"

	user2 := models.User{}
	user2.Email = "samael@work.com"
	user2.Password = "work"
	user2.Organization = "Xtern"
	user2.Role = "company"

	responseStatus, err := services.Register(ctx,user1)
	if responseStatus != http.StatusAccepted {
		t.Fail()
	}
	responseStatus, err = services.Register(ctx,user2)
	if responseStatus != http.StatusAccepted {
		t.Fail()
	}

	user1.Email = "xniccum@gmail.com"
	user1.Password = "admin1"
	user1.Organization = ""
	user1.Role = ""

	_, err = services.Login(ctx,user1)
	if err != nil {
		t.Fail()
	}

	user1.Email = "fake@gmail.com"
	_, err = services.Login(ctx,user1)
	if err == nil {
		t.Fail()
	}

	user1.Email = "xniccum@gmail.com"
	user1.Password = "admin999"
	_, err = services.Login(ctx,user1)
	if err == nil {
		t.Fail()
	}
}
