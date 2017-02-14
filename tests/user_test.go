package tests

import (
	"Xtern-Matching/models"
	"github.com/stretchr/testify/assert"
	"Xtern-Matching/handlers/services"
	"testing"
	"google.golang.org/appengine/aetest"
	"net/http"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

func GetUser1() models.User {
	user1 := models.User{}
	user1.Name = "Alex Crowley"
	user1.Email = "xniccum@gmail.com"
	user1.Password = "admin1"
	return user1
}

func GetUser2() models.User {
	user2 := models.User{}
	user2.Name = "Samuel Adams"
	user2.Email = "samael@work.com"
	user2.Password = "work"
	return user2
}

/*
	Returns responseStatus, userKey, orgKey, error
 */
func createUserAndOrg(ctx context.Context, user models.User) (int, *datastore.Key, *datastore.Key, error) {
	return createUserAndNamedOrg(ctx, user,"Dara Biosciences")
}

func createUserAndNamedOrg(ctx context.Context, user models.User, orgName string) (int, *datastore.Key, *datastore.Key, error) {
	orgKey, err := services.NewOrganization(ctx, "Dara Biosciences")
	if err != nil {
		return 500, nil, nil, err
	}
	time.Sleep(time.Millisecond * 500)
	responseStatus, userKey, err := services.Register(ctx, orgKey, user)
	if err != nil {
		return 0, nil, nil, err
	}
	return responseStatus, userKey, orgKey, nil
}

func TestRegister(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if !assert.Nil(t, err, "Error instantiating context") {
		t.Fatal(err)
	}
	defer done()
	user1 := GetUser1()
	user2 := GetUser2()


	responseStatus, _, orgKey, err := createUserAndOrg(ctx, user1)
	if !assert.Nil(t, err, "Error creating Student") {
		t.Fatal(err)
	}
	if !assert.Equal(t, responseStatus, http.StatusCreated, "Failed to register user1") {
		t.Fatal()
	}
	responseStatus, _, err = services.Register(ctx, orgKey, user2)
	if !assert.Equal(t, responseStatus, http.StatusCreated, "Failed to register user2") {
		t.Fatal()
	}

	responseStatus, _, err = services.Register(ctx, orgKey, user1)
	if !assert.Equal(t, responseStatus, http.StatusBadRequest, "User1 already exists") {
		t.Fatal()
	}
}

func TestLogin(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if !assert.Nil(t, err, "Error instantiating context") {
		t.Fatal(err)
	}
	defer done()

	user1 := GetUser1()
	user2 := GetUser2()
	orgKey, err := services.NewOrganization(ctx, "Dara Biosciences")
	time.Sleep(time.Millisecond * 500)
	if !assert.Nil(t, err, "Error creating Organization") {
		t.Fatal(err)
	}

	responseStatus, _, err := services.Register(ctx, orgKey, user1)
	if !assert.Equal(t, responseStatus, http.StatusCreated, "Failed to register user1") {
		t.Fatal()
	}
	responseStatus, _, err = services.Register(ctx, orgKey,user2)
	if !assert.Equal(t, responseStatus, http.StatusCreated, "Failed to register user2") {
		t.Fatal()
	}
	time.Sleep(time.Millisecond * 500)

	_, err = services.Login(ctx, user1.Email, user1.Password)
	if !assert.Nil(t, err, "Login should have validated") {
		t.Fail()
	}
	_, err = services.Login(ctx, user1.Email, user2.Password)
	if !assert.NotNil(t, err, "Login shouldn't have validated with user 2's password") {
		t.Fail()
	}
	user1.Email = "fake@gmail.com"
	_, err = services.Login(ctx, user1.Email, user1.Password)
	if !assert.NotNil(t, err, "Login shouldn't have validated with invalid email") {
		t.Fail()
	}
	user1.Email = "xniccum@gmail.com"
	user1.Password = "admin999"
	_, err = services.Login(ctx, user1.Email, user1.Password)
	if !assert.NotNil(t, err, "Login shouldn't have validated with invalid password") {
		t.Fail()
	}
}

func TestUserEmailPassUpdate(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if !assert.Nil(t, err, "Error instantiating context") {
		t.Fatal(err)
	}
	defer done()
	orgKey, err := services.NewOrganization(ctx, "Dara Biosciences")
	time.Sleep(time.Millisecond * 500)
	if !assert.Nil(t, err, "Error creating Organization") {
		t.Fatal(err)
	}
	user := GetUser1()
	responseStatus, userKey, err := services.Register(ctx, orgKey,user)
	if !assert.Equal(t, responseStatus, http.StatusCreated, "Failed to register user") {
		t.Fatal()
	}
	time.Sleep(time.Millisecond * 500)

	user.Password = "asdf"
	err = services.EditUser(ctx, userKey, "A.J.","xniccum@gmail.com", "asdf")
	if !assert.Nil(t, err, "Failed to update User") {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	_, err = services.Login(ctx, user.Email, user.Password)
	if !assert.Nil(t, err, "Login should have validated with changed password") {
		t.Fail()
	}

	user.Email = "go@test.com"
	user.Password = ""
	err = services.EditUser(ctx, userKey, user.Name, user.Email, user.Password)
	if !assert.Nil(t, err, "Failed to update User") {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	user.Password = "asdf"
	_, err = services.Login(ctx, user.Email, user.Password)
	if !assert.Nil(t, err, "Login should have validated with changed email") {
		t.Fail()
	}
}

/*
	tests GetUser and GetUsers
*/
func TestUserQueries(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if !assert.Nil(t, err, "Error instantiating context") {
		t.Fatal(err)
	}
	defer done()
	orgKey, err := services.NewOrganization(ctx, "Dara Biosciences")
	time.Sleep(time.Millisecond * 500)
	if !assert.Nil(t, err, "Error creating Organization") {
		t.Fatal(err)
	}
	user1 := GetUser1()
	user2 := GetUser2()

	responseStatus, userKey, err := services.Register(ctx, orgKey, user1)
	if !assert.Equal(t, responseStatus, http.StatusCreated, "Failed to register user1") {
		t.Fatal()
	}
	responseStatus, _, err = services.Register(ctx, orgKey, user2)
	if !assert.Equal(t, responseStatus, http.StatusCreated, "Failed to register user2") {
		t.Fatal()
	}
	time.Sleep(time.Millisecond * 500)
	users, _, err := services.GetUsers(ctx, orgKey)
	if !assert.Nil(t, err, "Failed to query Users on admin") {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	if !assert.Equal(t, 2, len(users), "User query expected to return 2 Users") {
		t.Fatal()
	}
	user, err := services.GetUser(ctx, userKey)
	if !assert.Nil(t, err, "Failed to query User by key") {
		t.Fatal(err)
	}
	//Because password gets hashed
	user1.Password = user.Password
	if !assert.Equal(t, user1, user, "Invalid User returned") {
		t.Fatal()
	}
}

func TestGetUserByOrgNae(t *testing.T)  {
	ctx, done, err := aetest.NewContext()
	if !assert.Nil(t, err, "Error instantiating context") {
		t.Fatal(err)
	}
	defer done()
	orgKey, err := services.NewOrganization(ctx, "Dara Biosciences")
	time.Sleep(time.Millisecond * 500)
	if !assert.Nil(t, err, "Error creating Organization") {
		t.Fatal(err)
	}
	user := GetUser1()
	responseStatus, userKey, err := services.Register(ctx, orgKey,user)
	if !assert.Equal(t, responseStatus, http.StatusCreated, "Failed to register user") {
		t.Fatal()
	}
	time.Sleep(time.Millisecond * 500)
	_, keys, err := services.GetUsersByOrgName(ctx, "Dara Biosciences")
	if !assert.Nil(t, err, "Error Running User query by organization") {
		t.Fatal(err)
	}
	if !assert.Equal(t, userKey, keys[0], "Key Mismatch in query result") {
		t.Fatal()
	}
}

func TestDeleteUser(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if !assert.Nil(t, err, "Error instantiating context") {
		t.Fatal(err)
	}
	defer done()
	orgKey, err := services.NewOrganization(ctx, "Dara Biosciences")
	time.Sleep(time.Millisecond * 500)
	if !assert.Nil(t, err, "Error creating Organization") {
		t.Fatal(err)
	}
	user := GetUser1()
	responseStatus, userKey, err := services.Register(ctx, orgKey,user)
	if !assert.Equal(t, responseStatus, http.StatusCreated, "Failed to register user") {
		t.Fatal()
	}
	time.Sleep(time.Millisecond * 500)
	err = services.DeleteUser(ctx, userKey)
	if !assert.Equal(t, responseStatus, http.StatusCreated, "Failed to Delete user") {
		t.Fatal()
	}
	time.Sleep(time.Millisecond * 500)

	user, err = services.GetUser(ctx, userKey)
	if !assert.NotNil(t, err, "GetUser should throw NoSuchEntity") {
		t.Fatal()
	}
}
