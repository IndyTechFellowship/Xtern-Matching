package services

import (
	"golang.org/x/net/context"
	"Xtern-Matching/models"
	"net/http"
	"google.golang.org/appengine/datastore"
)

func NewOrganization(ctx context.Context,name string) (*datastore.Key, error) {
	key := datastore.NewIncompleteKey(ctx, "Organization", nil)
	org := models.NewOrganization(name)
	key, err := datastore.Put(ctx, key, &org)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func GetOrganization(ctx context.Context, orgKey *datastore.Key) (models.Organization, error) {
	var org models.Organization
	err := datastore.Get(ctx, orgKey, &org)
	if err != nil {
		return models.Organization{}, err
	}
	return org, nil
}

func GetOrganizations(ctx context.Context) ([]models.Organization,[]*datastore.Key,error) {
	q := datastore.NewQuery("Organization").Project("Name")
	var orgs []models.Organization
	keys, err := q.GetAll(ctx, &orgs)
	if err != nil {
		return nil, nil, err
	}
	return orgs, keys, nil
}

func AddStudentToOrganization(ctx context.Context, orgKey *datastore.Key, studentKey *datastore.Key) (int64,error)  {
	//orgKey := datastore.NewKey(ctx, "Company", "", companyId, nil)
	var org models.Organization
	if err := datastore.Get(ctx, orgKey, &org); err != nil {
		return http.StatusInternalServerError, err
	}

	org.AddStudent(studentKey)

	if _, err := datastore.Put(ctx, orgKey, &org); err != nil {
		return http.StatusInternalServerError, err
	}
	return orgKey.IntID(), nil
}

func RemoveStudentFromOrganization(ctx context.Context, orgKey *datastore.Key, studentKey *datastore.Key) error  {
	//orgKey := datastore.NewKey(ctx, "Company", "", companyId, nil)
	var org models.Organization
	if err := datastore.Get(ctx, orgKey, &org); err != nil {
		return err
	}

	org.RemoveStudent(studentKey)

	if _, err := datastore.Put(ctx, orgKey, &org); err != nil {
		return err
	}
	return nil
}

// func MoveStudentInOrganization(ctx context.Context, orgKey int64, studentKey *datastore.Key, pos int) (int64,error)  {
// 	orgKey := datastore.NewKey(ctx, "Company", "", companyId, nil)
// 	var org models.Organization
// 	if err := datastore.Get(ctx, orgKey, &org); err != nil {
// 		return http.StatusInternalServerError, err
// 	}

// 	org.MoveStudent(studentKey, pos)

// 	if _, err := datastore.Put(ctx, orgKey, &org); err != nil {
// 		return http.StatusInternalServerError, err
// 	}
// 	return orgKey.IntID(), nil
// }

func SwitchStudentsInOrganization(ctx context.Context, orgKey *datastore.Key, student1Id *datastore.Key, student2Id *datastore.Key) (int64,error)  {
	var company models.Organization
	if err := datastore.Get(ctx, orgKey, &company); err != nil {
		return http.StatusInternalServerError, err
	}

	// company.Id = companyId
	if contains(company.Students, student1Id) && contains(company.Students, student2Id) {
		company.Students = switchElements(company.Students, student1Id, student2Id);
	} else {
		return http.StatusInternalServerError, nil
	}

	if _, err := datastore.Put(ctx, orgKey, &company); err != nil {
		return http.StatusInternalServerError, err
	}
	return orgKey.IntID(), nil
}

//func GetOrganization(ctx context.Context, orgKey datastore.Key) (models.Organization,error) {
//	//orgKey := datastore.NewKey(ctx, "Organization", "", _id, nil)
//	var org models.Organization
//	if err := datastore.Get(ctx, orgKey, &org); err != nil {
//		return models.Organization{}, err
//	}
//	return org, nil
//}

func switchElements(array []*datastore.Key, a *datastore.Key, b *datastore.Key) []*datastore.Key {
    for i := 0; i < len(array); i++ {
        if array[i] == a {
            array[i] = b
        } else if array[i] == b {
        	array[i] = a
        }
    }
    return array
}

func contains(array []*datastore.Key, element *datastore.Key) bool {
    for _, arrayElement := range array {
        if arrayElement == element {
    		return true
        }
    }
    return false
}