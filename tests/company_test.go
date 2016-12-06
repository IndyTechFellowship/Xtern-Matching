package tests

import (
	"testing"
	"google.golang.org/appengine/aetest"
	"Xtern-Matching/handlers/services"
	"Xtern-Matching/models"
	"time"
	"errors"
)

func TestCompany(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()
	
	//Need to be able to get company ID to test
	var student_1_id int64 = 1234
	var student_2_id int64 = 5678
	var company models.Company
	company.Name = "Darla Leach"
	//company.StudentIds []int64	= []int64{"Product Management","Software Engineer- Middle-tier Dev."}

	// Basic input
	_, err = services.NewCompany(ctx, company)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	
	companies, err := services.GetCompanies(ctx)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	
	company_id := companies[0].Id
	
	_, err = services.AddStudentIdToCompanyList(ctx, company_id, student_1_id)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	
	_, err = services.AddStudentIdToCompanyList(ctx, company_id, student_2_id)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	
	_, err = services.SwitchStudentIdsInCompanyList(ctx, company_id, student_1_id, student_2_id)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	
	test_company, err := services.GetCompany(ctx, company_id)
	if err != nil {
		t.Fatal(err)
	}	
	time.Sleep(time.Millisecond * 500)
	
	if test_company.StudentIds[0] != student_2_id || test_company.StudentIds[1] != student_1_id {
		t.Fatal(errors.New("Student ID mismatch"))
	}
	
	_, err = services.RemoveStudentIdFromCompanyList(ctx, company_id, student_2_id)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)
	
	test_company, err = services.GetCompany(ctx, company_id)
	if err != nil {
		t.Fatal(err)
	}	
	time.Sleep(time.Millisecond * 500)
	
	if test_company.StudentIds[0] != student_1_id {
		t.Fatal(errors.New("Student ID mismatch after removal"))
	}
}
