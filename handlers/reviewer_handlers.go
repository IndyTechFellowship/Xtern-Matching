package handlers

import (
	"net/http"
	"google.golang.org/appengine"
	"encoding/json"
	"Xtern-Matching/handlers/services"
	"log"
	"google.golang.org/appengine/datastore"
	"math/rand"
	"math"
	"Xtern-Matching/models"
)

func CreateReviewGroups(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	minStudents := int(dat["minStudents"].(float64))
	minReviewers := int(dat["minReviewers"].(float64))

	//: Wipe existing groups
	_, oldReviewGroupKeys, err := services.GetReviewGroups(ctx, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	} else {
		err := datastore.DeleteMulti(ctx, oldReviewGroupKeys)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
	}

	//: get array of keys for reviewers
	_, reviewerKeys, err := services.GetUsersByOrgName(ctx, "Reviewers")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	//: get array of keys for students
	_, studentKeys, err := services.GetStudents(ctx, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// log.Println(reviewerKeys)

	// shuffle the arrays
	for i := range reviewerKeys {
    	j := rand.Intn(i + 1)
    	reviewerKeys[i], reviewerKeys[j] = reviewerKeys[j], reviewerKeys[i]
	}

	for i := range studentKeys {
    	j := rand.Intn(i + 1)
    	studentKeys[i], studentKeys[j] = studentKeys[j], studentKeys[i]
	}

	//: loop: create and assign reviewers/students to groups - limit max groups to min number of students or reviewers
	maxGroups := int(math.Min(float64(len(studentKeys)/minStudents), float64(len(reviewerKeys)/minReviewers)));
	// maxGroups = 2

	// compute num groups from input group size
	reviewGroups := make([]models.ReviewGroup,maxGroups)
	reviewGroupKeys := make([]*datastore.Key,maxGroups)

	var reviewerArrayIndex = 0;
	var studentArrayIndex = 0;
	var numReviewersPerGroup = int(len(reviewerKeys)/maxGroups);
	var numStudentsPerGroup = int(len(studentKeys)/maxGroups);

	log.Println("numReviewersPerGroup: ", numReviewersPerGroup)
	log.Println("numStudentsPerGroup: ", numStudentsPerGroup)
	log.Println("maxGroups: ", maxGroups)


	for i := 0; i < maxGroups; i++ {
		var newReviewGroup models.ReviewGroup
		
		reviewGroupKeys = append(reviewGroupKeys, datastore.NewIncompleteKey(ctx, "ReviewGroup", nil))

		for j := 0; j < numReviewersPerGroup; j++ {
			newReviewGroup.Reviewers = append(newReviewGroup.Reviewers, reviewerKeys[reviewerArrayIndex + j])
		}
		reviewerArrayIndex += numReviewersPerGroup

		for k := 0; k < numStudentsPerGroup; k++ {
			newReviewGroup.Students = append(newReviewGroup.Students, studentKeys[studentArrayIndex + k])
		}
		studentArrayIndex += numStudentsPerGroup
		reviewGroups = append(reviewGroups, newReviewGroup)
		// log.Println("new group made")
	}

	//TODO: assign remaining students and reviewers while looping groups one at a time

	var reviewGroupIndex = 0;

	if(reviewerArrayIndex < len(reviewerKeys)) {
		for reviewerArrayIndex < len(reviewerKeys) {
			reviewGroups[reviewGroupIndex].Reviewers = append(reviewGroups[reviewGroupIndex].Reviewers, reviewerKeys[reviewerArrayIndex])
			reviewerArrayIndex++
			reviewGroupIndex++
			if(reviewGroupIndex >= len(reviewGroups)) {
				reviewGroupIndex = 0;
			}
			log.Println("extra reviewer added")
		}
		reviewGroupIndex = 0;
	}

	if(studentArrayIndex < len(studentKeys)) {
		for studentArrayIndex < len(studentKeys) {
			reviewGroups[reviewGroupIndex].Students = append(reviewGroups[reviewGroupIndex].Students, studentKeys[studentArrayIndex])
			studentArrayIndex++
			reviewGroupIndex++
			if(reviewGroupIndex >= len(reviewGroups)) {
				reviewGroupIndex = 0;
			}
			log.Println("extra reviewer added")
		}
	}

	// _, err = datastore.PutMulti(ctx, reviewGroupKeys, &reviewGroups)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	http.Error(w, err.Error(), 500)
	// }

	for i := 0; i < len(reviewGroups); i ++ {
		log.Println(reviewGroups[i]);
		datastore.Put(ctx, reviewGroupKeys[i], &reviewGroups[i])
	}
	w.WriteHeader(http.StatusOK)
}

func GetReviewGroups(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	reviewGroups, reviewGroupKeys, err := services.GetReviewGroups(ctx, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	type Response struct {
		Keys []*datastore.Key		`json:"keys"`
		ReviewGroups []models.ReviewGroup		`json:"users"`
	}
	response := Response{Keys: reviewGroupKeys, ReviewGroups: reviewGroups}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetReviewGroupForReviewer(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	reviewerKey, err := datastore.DecodeKey(dat["reviewerKey"].(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	reviewGroup, reviewGroupKey, err := services.GetReviewGroupForReviewer(ctx, reviewerKey)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}


	type Response struct {
		Key *datastore.Key		`json:"keys"`
		ReviewGroup models.ReviewGroup		`json:"users"`
	}
	response := Response{Key: reviewGroupKey, ReviewGroup: reviewGroup}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}