package services

import (
	"golang.org/x/net/context"
	"Xtern-Matching/models"
	"net/http"
	"google.golang.org/appengine/datastore"
)

func removeId(ids []int64, idToRemove int64) []int64 {
    filteredIds := ids[:0]
    for _, id := range ids {
        if id != idToRemove {
            filteredIds = append(filteredIds, id)
        }
    }
    return filteredIds
}

func contains(array []int64, element int64) bool {
    for _, arrayElement := range array {
        if arrayElement == element {
    		return true
        }
    }
    return false
}

func switchElements(array []int64, a int64, b int64) []int64 {
    for i := 0; i < len(array); i++ {
        if array[i] == a {
            array[i] = b
        } else if array[i] == b {
        	array[i] = a
        }
    }
    return array
}

func AddStudentIdToCompanyList(ctx context.Context,companyId int64, studentId int64) (int64,error)  {
	companyKey := datastore.NewKey(ctx, "Company", "", companyId, nil)
	var company models.Company
	if err := datastore.Get(ctx, companyKey, &company); err != nil {
		return http.StatusInternalServerError, err
	}

	company.Id = companyId
	if !contains(company.StudentIds, studentId) {
		company.StudentIds = append(company.StudentIds, studentId);
	}

	if _, err := datastore.Put(ctx, companyKey, &company); err != nil {
		return http.StatusInternalServerError, err
	}
	return company.Id, nil
}

func RemoveStudentIdFromCompanyList(ctx context.Context,companyId int64, studentId int64) (int64,error)  {
	companyKey := datastore.NewKey(ctx, "Company", "", companyId, nil)
	var company models.Company
	if err := datastore.Get(ctx, companyKey, &company); err != nil {
		return http.StatusInternalServerError, err
	}

	company.Id = companyId
	company.StudentIds = removeId(company.StudentIds, studentId);
	// company.StudentIds = company.StudentIds.remove(studentId);

	if _, err := datastore.Put(ctx, companyKey, &company); err != nil {
		return http.StatusInternalServerError, err
	}
	return company.Id, nil
}

func SwitchStudentIdsInCompanyList(ctx context.Context,companyId int64, student1Id int64, student2Id int64) (int64,error)  {
	companyKey := datastore.NewKey(ctx, "Company", "", companyId, nil)
	var company models.Company
	if err := datastore.Get(ctx, companyKey, &company); err != nil {
		return http.StatusInternalServerError, err
	}

	company.Id = companyId
	if contains(company.StudentIds, student1Id) && contains(company.StudentIds, student2Id) {
		company.StudentIds = switchElements(company.StudentIds, student1Id, student2Id);
	}

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