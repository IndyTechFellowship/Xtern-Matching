package handlers

import (
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"Xtern-Matching/handlers/services"
	"Xtern-Matching/models"
	"google.golang.org/appengine/datastore"
)

func StartUp(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "startup succesful")
}

func WarmUp(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	query := datastore.NewQuery("User")
	count, _ := query.Count(ctx)
	if count == 0 {
		//TODO eventually change this to load from yml or json
		//Seed Database
		var users [4]models.User
		users[0] = models.User{
			Name: "Austin Niccum",
			Email: "xniccum@gmail.com",
			Password: "admin1",
			Organization: "Techpoint",
			Role: "Admin"}
		users[1] = models.User{
			Name: "Steven Doolan",
			Email: "srdoolan3@gmail.com",
			Password: "admin1",
			Organization: "Techpoint",
			Role: "Admin"}
		users[2] = models.User{
			Name: "Davis Nygren",
			Email: "DavisNygren@gmail.com",
			Password: "admin1",
			Organization: "Techpoint",
			Role: "Admin"}
		users[3] = models.User{
			Name: "Alex Crowley",
			Email: "acrow94@gmail.com",
			Password: "admin1",
			Organization: "Techpoint",
			Role: "Admin"}
		for _, user := range users {
			_, err := services.Register(ctx,user)
			if err != nil {
				log.Debugf(ctx, "Error in Seeding")
				return
			}
		}

	}

	log.Infof(ctx, "warmup done")
}
