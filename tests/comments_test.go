package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"google.golang.org/appengine/aetest"
	"Xtern-Matching/handlers/services"
	"time"
)

func TestCommentPost(t *testing.T) {
	t.Parallel()
	ctx, done, err := aetest.NewContext()
	if !assert.Nil(t, err, "Error instantiating context") {
		t.Fatal(err)
	}
	defer done()
	_, studentKey, err := createStudent(ctx)
	if !assert.Nil(t, err, "Error creating student") {
		t.Fatal(err)
	}
	user := GetUser1()
	_, userKey, _, err := createUserAndOrg(ctx, user)
	if !assert.Nil(t, err, "Error creating user") {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	_, _, err = services.AddComment(ctx, studentKey,"Good applicant", user.Name, userKey)
	if !assert.Nil(t, err, "Error adding comment") {
		t.Fatal(err)
	}
}

func TestCommentGet(t *testing.T) {
	t.Parallel()
	ctx, done, err := aetest.NewContext()
	if !assert.Nil(t, err, "Error instantiating context") {
		t.Fatal(err)
	}
	defer done()
	_, studentKey, err := createStudent(ctx)
	if !assert.Nil(t, err, "Error creating student") {
		t.Fatal(err)
	}
	user := GetUser1()
	user2 := GetUser2()
	_, userKey, orgKey1, err := createUserAndOrg(ctx, user)
	if !assert.Nil(t, err, "Error creating user") {
		t.Fatal(err)
	}
	_, userKey2, err := services.Register(ctx, orgKey1, user2)
	if !assert.Nil(t, err, "Error creating user") {
		t.Fatal(err)
	}
	_, userKey3, orgKey2, err := createUserAndNamedOrg(ctx, user2, "AMD")
	if !assert.Nil(t, err, "Error creating user") {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	_, _, err = services.AddComment(ctx, studentKey,"Good applicant", user.Name, userKey)
	if !assert.Nil(t, err, "Error adding comment") {
		t.Fatal(err)
	}
	_, _, err = services.AddComment(ctx, studentKey,"Good applicant", user.Name, userKey2)
	if !assert.Nil(t, err, "Error adding comment") {
		t.Fatal(err)
	}
	_, _, err = services.AddComment(ctx, studentKey,"Good applicant", user.Name, userKey3)
	if !assert.Nil(t, err, "Error adding comment") {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	comments, _, err := services.GetComments(ctx, studentKey, orgKey1)
	if !assert.Nil(t, err, "Error getting comments") {
		t.Fatal(err)
	}
	if !assert.Equal(t,2, len(comments), "Unexpected comment size") {
		t.Fatal(err)
	}
	comments, _, err = services.GetComments(ctx, studentKey, orgKey2)
	if !assert.Nil(t, err, "Error getting comments") {
		t.Fatal(err)
	}
	if !assert.Equal(t,1, len(comments), "Unexpected comment size") {
		t.Fatal(err)
	}
}

func TestCommentEdit(t *testing.T)  {
	t.Parallel()
	ctx, done, err := aetest.NewContext()
	if !assert.Nil(t, err, "Error instantiating context") {
		t.Fatal(err)
	}
	defer done()
	_, studentKey, err := createStudent(ctx)
	if !assert.Nil(t, err, "Error creating student") {
		t.Fatal(err)
	}
	user := GetUser1()
	_, userKey, orgKey, err := createUserAndOrg(ctx, user)
	if !assert.Nil(t, err, "Error creating user") {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	_, commentKey, err := services.AddComment(ctx, studentKey,"Good applicant", user.Name, userKey)
	if !assert.Nil(t, err, "Error adding comment") {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	_, err = services.EditComment(ctx, commentKey, "We need this candidate")
	if !assert.Nil(t, err, "Error editing comment") {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	comments, _, err := services.GetComments(ctx, studentKey, orgKey)
	if !assert.Nil(t, err, "Error getting comments") {
		t.Fatal(err)
	}
	if !assert.Equal(t,"We need this candidate", comments[0].Message, "Persistence issue in updating conmments") {
		t.Fatal(err)
	}

}

func TestCommentDelete(t *testing.T)  {
	t.Parallel()
	ctx, done, err := aetest.NewContext()
	if !assert.Nil(t, err, "Error instantiating context") {
		t.Fatal(err)
	}
	defer done()
	_, studentKey, err := createStudent(ctx)
	if !assert.Nil(t, err, "Error creating student") {
		t.Fatal(err)
	}
	user := GetUser1()
	_, userKey, orgKey, err := createUserAndOrg(ctx, user)
	if !assert.Nil(t, err, "Error creating user") {
		t.Fatal(err)
	}
	_, userKey2, err := services.Register(ctx, orgKey, GetUser2())
	if !assert.Nil(t, err, "Error creating user") {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	_, commentKey, err := services.AddComment(ctx, studentKey,"Good applicant", user.Name, userKey)
	if !assert.Nil(t, err, "Error adding comment") {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	err = services.DeleteComment(ctx, commentKey, userKey2)
	if !assert.NotNil(t, err, "Should have thrown error when deleting with wrong user") {
		t.Fatal(err)
	}
	err = services.DeleteComment(ctx, commentKey, userKey)
	if !assert.Nil(t, err, "Error occured when deleting user") {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	comments, _, err := services.GetComments(ctx, studentKey, orgKey)
	if !assert.Nil(t, err, "Error getting comments") {
		t.Fatal(err)
	}
	if !assert.Equal(t,0, len(comments), "Unexpected comment size") {
		t.Fatal(err)
	}
}