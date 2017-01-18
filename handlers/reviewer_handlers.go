package handlers

import (
	"net/http"
	"google.golang.org/appengine"
	"encoding/json"
	"Xtern-Matching/handlers/services"
	"log"
	"google.golang.org/appengine/datastore"
)

func CreateReviewGroups(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// minStudents := dat["minStudents"].(int)
	// minReviewers := dat["minReviewers"].(int)

	//TODO: Wipe existing groups
	_, ReviewGroupKeys, err := services.GetReviewGroups(ctx, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	} else {
		err := datastore.DeleteMulti(ctx, ReviewGroupKeys)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
	}

	//TODO: get array of keys for reviewers

	//TODO: get array of keys for students

	//TODO: loop: create and assign reviewers/students to groups - limit max groups to min number of students or reviewers

	//TODO: assign remaining students and reviewers while looping groups one at a time










	// user := context.Get(r, "user")
	// mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	// orgKey, err := datastore.DecodeKey(mapClaims["org"].(string))
	// if err != nil {
	// 	log.Println(err.Error())
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }
	// position :=  int(dat["position"].(float64));

	// _, err = services.MoveStudentInOrganization(ctx, orgKey, studentKey, position)
	// if err != nil {
	// 	log.Print(err)
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }
	w.WriteHeader(http.StatusOK)

}