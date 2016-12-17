package tests
//
//import (
//	"Xtern-Matching/handlers/services"
//	"Xtern-Matching/models"
//	"net/http"
//	"testing"
//	"time"
//
//	"github.com/stretchr/testify/assert"
//	"google.golang.org/appengine/aetest"
//)
//
//func GetUser1() models.User {
//	user1 := models.User{}
//	user1.Email = "xniccum@gmail.com"
//	user1.Password = "admin1"
//	user1.Organization = "Xtern"
//	user1.Role = "admin"
//	return user1
//}
//
//func GetUser2() models.User {
//	user2 := models.User{}
//	user2.Email = "samael@work.com"
//	user2.Password = "work"
//	user2.Organization = "Xtern"
//	user2.Role = "company"
//	return user2
//}
//
//func TestRegister(t *testing.T) {
//	ctx, done, err := aetest.NewContext()
//	if !assert.Nil(t, err, "Error instantiating context") {
//		t.Fatal(err)
//	}
//	defer done()
//	user1 := GetUser1()
//	user2 := GetUser2()
//
//	responseStatus, err := services.Register(ctx, user1)
//	if !assert.Equal(t, responseStatus, http.StatusCreated, "Failed to register user1") {
//		t.Fatal()
//	}
//	responseStatus, err = services.Register(ctx, user2)
//	if !assert.Equal(t, responseStatus, http.StatusCreated, "Failed to register user2") {
//		t.Fatal()
//	}
//
//	responseStatus, err = services.Register(ctx, user1)
//	if !assert.Equal(t, responseStatus, http.StatusBadRequest, "User1 already exists") {
//		t.Fatal()
//	}
//}
//
//func TestLogin(t *testing.T) {
//	ctx, done, err := aetest.NewContext()
//	if !assert.Nil(t, err, "Error instantiating context") {
//		t.Fatal(err)
//	}
//	defer done()
//
//	user1 := GetUser1()
//
//	user2 := GetUser2()
//
//	responseStatus, err := services.Register(ctx, user1)
//	if !assert.Equal(t, responseStatus, http.StatusCreated, "Failed to register user1") {
//		t.Fatal()
//	}
//
//	responseStatus, err = services.Register(ctx, user2)
//
//	if !assert.Equal(t, responseStatus, http.StatusCreated, "Failed to register user2") {
//		t.Fatal()
//	}
//
//	user1.Email = "xniccum@gmail.com"
//	user1.Password = "admin1"
//	user1.Organization = ""
//	user1.Role = ""
//
//	//TODO: This shouldn't pass
//	_, err = services.Login(ctx, user1)
//	if !assert.Nil(t, err, "Login shouldn't have validated with missing role and organization") {
//		t.Fail()
//	}
//
//	user1.Email = "fake@gmail.com"
//	_, err = services.Login(ctx, user1)
//	if !assert.NotNil(t, err, "Login shouldn't have validated with invalid email") {
//		t.Fail()
//	}
//
//	user1.Email = "xniccum@gmail.com"
//	user1.Password = "admin999"
//	_, err = services.Login(ctx, user1)
//	if !assert.NotNil(t, err, "Login shouldn't have validated with invalid password") {
//		t.Fail()
//	}
//}
//
//func TestUserEmailPassUpdate(t *testing.T) {
//	ctx, done, err := aetest.NewContext()
//	if !assert.Nil(t, err, "Error instantiating context") {
//		t.Fatal(err)
//	}
//	defer done()
//	user := GetUser1()
//	responseStatus, err := services.Register(ctx, user)
//	if !assert.Equal(t, responseStatus, http.StatusCreated, "Failed to register user") {
//		t.Fatal()
//	}
//
//	user.Password = "asdf"
//	err = services.UpdateUser(ctx, &user)
//	if !assert.Nil(t, err, "Failed to update User") {
//		t.Fatal(err)
//	}
//	time.Sleep(time.Millisecond * 500)
//	user.Password = "asdf"
//	_, err = services.Login(ctx, user)
//	if !assert.Nil(t, err, "Login should have validated with changed password") {
//		t.Fail()
//	}
//
//	//Not updating the password this time
//	users, err := services.GetUsers(ctx, "Xtern", "admin")
//	time.Sleep(time.Millisecond * 500)
//	user = users[0]
//	user.Email = "go@test.com"
//	user.Password = "********"
//	err = services.UpdateUser(ctx, &user)
//	if !assert.Nil(t, err, "Failed to update User") {
//		t.Fatal(err)
//	}
//	time.Sleep(time.Millisecond * 500)
//	user.Password = "asdf"
//	_, err = services.Login(ctx, user)
//	if !assert.Nil(t, err, "Login should have validated with changed email") {
//		t.Fail()
//	}
//}
//
///*
//	tests GetUser and GetUsers
//*/
//func TestUserQueries(t *testing.T) {
//	ctx, done, err := aetest.NewContext()
//	if !assert.Nil(t, err, "Error instantiating context") {
//		t.Fatal(err)
//	}
//	defer done()
//	user1 := GetUser1()
//	user2 := GetUser2()
//	user2.Role = "admin"
//
//	responseStatus, err := services.Register(ctx, user1)
//	if !assert.Equal(t, responseStatus, http.StatusCreated, "Failed to register user1") {
//		t.Fatal()
//	}
//	responseStatus, err = services.Register(ctx, user2)
//	if !assert.Equal(t, responseStatus, http.StatusCreated, "Failed to register user2") {
//		t.Fatal()
//	}
//	time.Sleep(time.Millisecond * 500)
//	users, err := services.GetUsers(ctx, "Xtern", "admin")
//	if !assert.Nil(t, err, "Failed to query Users on admin") {
//		t.Fatal(err)
//	}
//	time.Sleep(time.Millisecond * 500)
//	if !assert.Equal(t, 2, len(users), "User query expected to return 2 Users") {
//		t.Fatal()
//	}
//	user, err := services.GetUser(ctx, users[0].Id)
//	if !assert.Nil(t, err, "Failed to query Users on admin") {
//		t.Fatal(err)
//	}
//	if !assert.Equal(t, users[0], user, "Invalid User returned") {
//		t.Fatal()
//	}
//}
//
//func TestDeleteUser(t *testing.T) {
//	ctx, done, err := aetest.NewContext()
//	if !assert.Nil(t, err, "Error instantiating context") {
//		t.Fatal(err)
//	}
//	defer done()
//	user := GetUser1()
//	responseStatus, err := services.Register(ctx, user)
//	if !assert.Equal(t, responseStatus, http.StatusCreated, "Failed to register user") {
//		t.Fatal()
//	}
//	time.Sleep(time.Millisecond * 500)
//	users, err := services.GetUsers(ctx, "Xtern", "admin")
//	time.Sleep(time.Millisecond * 500)
//	user = users[0]
//	services.DeleteUser(ctx, user.Id)
//	time.Sleep(time.Millisecond * 500)
//	users, err = services.GetUsers(ctx, "Xtern", "admin")
//	if !assert.Equal(t, 0, len(users), "Expected no users after deletion") {
//		t.Fatal()
//	}
//}
