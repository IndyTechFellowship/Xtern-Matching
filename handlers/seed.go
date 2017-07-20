package handlers


import (
	"net/http"
	"google.golang.org/appengine"
	"Xtern-Matching/handlers/services"
	"Xtern-Matching/models"
	"google.golang.org/appengine/datastore"
	"os"
	"encoding/json"
	"google.golang.org/appengine/log"
	"golang.org/x/net/context"
)

func SeedStudents(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	//Temporary
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS","environments/development/cloudstore-dev.json")

	//Seed Database
	type Seeds struct {
		Organizations []map[string]string 	`json:"organizations"`
		Users         []map[string]string       `json:"users"`
		Students      []models.Student          `json:"students"`
	}
	var seeds Seeds

	seedFile, err := os.Open("environments/development/seed.json")
	if err != nil {
		log.Errorf(ctx, "Error: Problem reading seed")
		return
	}
	defer seedFile.Close()
	jsonParser := json.NewDecoder(seedFile)
	if err = jsonParser.Decode(&seeds); err != nil {
		log.Errorf(ctx, "Error: Problem decoding seed file")
		return
	}

	query := datastore.NewQuery("Student")
	count, _ := query.Count(ctx)
	if count == 0 {
		log.Infof(ctx, "Seeding Students")
		for _, student := range seeds.Students {
			//TODO make seed details correctly
			if _, err = services.NewStudent(ctx, student); err != nil {
				log.Errorf(ctx, "Problem seeding students: " + err.Error())
				return
			}
		}
	}
	log.Infof(ctx, "Seed Students Done")
}



func SeedOrgs(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	//Temporary
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS","environments/development/cloudstore-dev.json")

	//Seed Database
	type Seeds struct {
		Organizations []map[string]string 	`json:"organizations"`
		Users         []map[string]string       `json:"users"`
		Students      []models.Student          `json:"students"`
	}
	var seeds Seeds

	seedFile, err := os.Open("environments/development/seed.json")
	if err != nil {
		log.Errorf(ctx, "Error: Problem reading seed")
		return
	}
	defer seedFile.Close()
	jsonParser := json.NewDecoder(seedFile)
	if err = jsonParser.Decode(&seeds); err != nil {
		log.Errorf(ctx, "Error: Problem decoding seed file")
		return
	}

	query := datastore.NewQuery("User")
	count, _ := query.Count(ctx)
	if count == 0 {
		log.Infof(ctx, "Seeding Organizations")
		orgs := map[string]*datastore.Key{}
		err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
			for _, org := range seeds.Organizations {
				key, err := services.NewOrganization(ctx, org["name"])
				if err != nil {
					return err
				}
				orgs[org["name"]] = key
			}
			return nil
		}, &datastore.TransactionOptions{XG:true})
		if err != nil {
			log.Errorf(ctx, "Problem seeding organization: " + err.Error())
			return
		} else {
			log.Infof(ctx, "Seeding Users")
			err = datastore.RunInTransaction(ctx, func(ctx context.Context) error {
				for _, userMap := range seeds.Users {
					var user models.User
					user.Name = userMap["name"]
					user.Email = userMap["email"]
					user.Password = userMap["password"]
					if _, err = services.Register(ctx, orgs[userMap["org"]], user); err != nil {
						return err
					}
				}
				return nil
			}, &datastore.TransactionOptions{XG:true})
			if err != nil {
				log.Errorf(ctx, "Problem seeding organization: " + err.Error())
				return
			}
		}
	}
	log.Infof(ctx, "Seed Orgs Done")
}
