package tests

import (
	"github.com/stretchr/testify/assert"
	"Xtern-Matching/handlers/services"
	"testing"
	"google.golang.org/appengine/aetest"
	"time"
)

func TestOrganizationPost(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if !assert.Nil(t, err, "Error instantiating context") {
		t.Fatal(err)
	}
	defer done()

	_, err = services.NewOrganization(ctx, "Dara Biosciences")
	if !assert.Nil(t, err, "Error creating Organization") {
		t.Fatal(err)
	}

}

func TestOrganizationGet(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if !assert.Nil(t, err, "Error instantiating context") {
		t.Fatal(err)
	}
	defer done()

	key, err := services.NewOrganization(ctx, "Dara Biosciences")
	time.Sleep(time.Millisecond * 500)
	if !assert.Nil(t, err, "Error creating Organization") {
		t.Fatal(err)
	}

	_, keys, err := services.GetOrganizations(ctx)
	if !assert.Nil(t, err, "Error retrieving Organizations") {
		t.Fatal(err)
	}
	if !assert.Equal(t, key, keys[0],"Error, Key Mismatch in retrieval") {
		t.Fatal(err)
	}
	_, err = services.GetOrganization(ctx, key)
	if !assert.Nil(t, err, "Error retrieving Organization") {
		t.Fatal(err)
	}
}

/*
	TODO: Uncomment when organization switch call is fixed
func TestOrganizationPreferenceList(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	student := GetStudent1()
	student_1_id, err := services.NewStudent(ctx, student)
	student_2_id, err := services.NewStudent(ctx, student)
	// Basic input
	company_id, err := services.NewOrganization(ctx, "Dara Biosciences")
	if !assert.Nil(t, err, "Error creating Company") {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)

	_, err = services.AddStudentToOrganization(ctx, company_id, student_1_id)
	if !assert.Nil(t, err, "Error adding to Company list") {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)

	_, err = services.AddStudentToOrganization(ctx, company_id, student_2_id)
	if !assert.Nil(t, err, "Error adding to Company list") {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)

	_, err = services.SwitchStudentsInOrganization(ctx, company_id, student_1_id, student_2_id)
	if !assert.Nil(t, err, "Error switching students in Company list") {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)

	test_company, err := services.GetOrganization(ctx, company_id)
	if !assert.Nil(t, err, "Error retrieving Company") {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	if !assert.Equal(t, student_2_id, test_company.Students[0], "Student 2 id mismatch") ||
		!assert.Equal(t, student_1_id, test_company.Students[1], "Student 1 id mismatch") {
		t.Fatal()
	}

	err = services.RemoveStudentFromOrganization(ctx, company_id, student_2_id)
	if !assert.Nil(t, err, "Error removing from Organization list") {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)

	test_company, err = services.GetOrganization(ctx, company_id)
	if !assert.Nil(t, err, "Error retrieving Company") {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)

	if !assert.Equal(t, student_1_id, test_company.Students[0], "Student 1 id mismatch after deletion") {
		t.Fatal()
	}
}
*/