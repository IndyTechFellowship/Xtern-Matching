package tests
//
//import (
//	"Xtern-Matching/handlers/services"
//	"Xtern-Matching/models"
//	"testing"
//	"time"
//
//	"github.com/stretchr/testify/assert"
//	"google.golang.org/appengine/aetest"
//)
//
//func TestCompanyPost(t *testing.T) {
//	ctx, done, err := aetest.NewContext()
//	if !assert.Nil(t, err, "Error instantiating context") {
//		t.Fatal(err)
//	}
//	defer done()
//
//	var company models.Company
//	company.Name = "Darla Leach"
//	_, err = services.NewCompany(ctx, company)
//	if !assert.Nil(t, err, "Error creating Company") {
//		t.Fatal(err)
//	}
//
//}
//
//func TestCompanyGet(t *testing.T) {
//	ctx, done, err := aetest.NewContext()
//	if !assert.Nil(t, err, "Error instantiating context") {
//		t.Fatal(err)
//	}
//	defer done()
//
//	var company models.Company
//	company.Name = "Darla Leach"
//	_, err = services.NewCompany(ctx, company)
//	if !assert.Nil(t, err, "Error creating Company") {
//		t.Fatal(err)
//	}
//	_, err = services.GetCompanies(ctx)
//	if !assert.Nil(t, err, "Error retrieving Companies") {
//		t.Fatal(err)
//	}
//}
//
//func TestCompanyPreferenceList(t *testing.T) {
//	ctx, done, err := aetest.NewContext()
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer done()
//
//	//Need to be able to get company ID to test
//	var student_1_id int64 = 1234
//	var student_2_id int64 = 5678
//	var company models.Company
//	company.Name = "Darla Leach"
//	//company.StudentIds []int64	= []int64{"Product Management","Software Engineer- Middle-tier Dev."}
//
//	// Basic input
//	_, err = services.NewCompany(ctx, company)
//	if !assert.Nil(t, err, "Error creating Company") {
//		t.Fatal(err)
//	}
//	time.Sleep(time.Millisecond * 500)
//
//	companies, err := services.GetCompanies(ctx)
//	if !assert.Nil(t, err, "Error retrieving Companies") {
//		t.Fatal(err)
//	}
//	time.Sleep(time.Millisecond * 500)
//
//	company_id := companies[0].Id
//
//	_, err = services.AddStudentIdToCompanyList(ctx, company_id, student_1_id)
//	if !assert.Nil(t, err, "Error adding to Company list") {
//		t.Fatal(err)
//	}
//	time.Sleep(time.Millisecond * 500)
//
//	_, err = services.AddStudentIdToCompanyList(ctx, company_id, student_2_id)
//	if !assert.Nil(t, err, "Error adding to Company list") {
//		t.Fatal(err)
//	}
//	time.Sleep(time.Millisecond * 500)
//
//	_, err = services.SwitchStudentIdsInCompanyList(ctx, company_id, student_1_id, student_2_id)
//	if !assert.Nil(t, err, "Error switching students in Company list") {
//		t.Fatal(err)
//	}
//	time.Sleep(time.Millisecond * 500)
//
//	test_company, err := services.GetCompany(ctx, company_id)
//	if !assert.Nil(t, err, "Error retrieving Company") {
//		t.Fatal(err)
//	}
//	time.Sleep(time.Millisecond * 500)
//	if !assert.Equal(t, student_2_id, test_company.StudentIds[0], "Student 2 id mismatch") ||
//		!assert.Equal(t, student_1_id, test_company.StudentIds[1], "Student 1 id mismatch") {
//		t.Fatal()
//	}
//
//	_, err = services.RemoveStudentIdFromCompanyList(ctx, company_id, student_2_id)
//	if !assert.Nil(t, err, "Error removing from Company list") {
//		t.Fatal(err)
//	}
//	time.Sleep(time.Millisecond * 500)
//
//	test_company, err = services.GetCompany(ctx, company_id)
//	if !assert.Nil(t, err, "Error retrieving Company") {
//		t.Fatal(err)
//	}
//	time.Sleep(time.Millisecond * 500)
//
//	if !assert.Equal(t, student_1_id, test_company.StudentIds[0], "Student 1 id mismatch after deletion") {
//		t.Fatal()
//	}
//}
