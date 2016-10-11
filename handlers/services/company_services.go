package services

import (
	"golang.org/x/net/context"
	"Xtern-Matching/models"
	"net/http"
	"google.golang.org/appengine/datastore"
)

func AddStudentIdToCompanyList(ctx context.Context,companyId int64, studentId int64) (int64,error)  {
	companyKey := datastore.NewKey(ctx, "Company", "", companyId, nil)
	var company models.Company
	if err := datastore.Get(ctx, companyKey, &company); err != nil {
		return http.StatusInternalServerError, err
	}

	company.Id = companyId
	company.StudentIds = append(company.StudentIds, studentId);

	if _, err := datastore.Put(ctx, companyKey, &company); err != nil {
		return http.StatusInternalServerError, err
	}
	return company.Id, nil
}

func NewCompany(ctx context.Context,company models.Company) (int,error) {
	key := datastore.NewIncompleteKey(ctx, "Company", nil)
	if _, err := datastore.Put(ctx, key, &company); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusAccepted, nil
}

func GetCompany(ctx context.Context,_id int64) (models.Company,error) {
	companyKey := datastore.NewKey(ctx, "Company", "", _id, nil)
	var company models.Company
	if err := datastore.Get(ctx, companyKey, &company); err != nil {
		return models.Company{}, err
	}
	company.Id = companyKey.IntID()
	return company, nil
}