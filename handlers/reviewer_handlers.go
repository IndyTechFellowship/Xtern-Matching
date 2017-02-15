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
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
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

func GetCurrentProgress(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)
	var noReviews, someReviews, allReviews int = 0,0,0

	reviewGroups, _, err := services.GetReviewGroups(ctx, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	for i:=0; i<len(reviewGroups); i++ {
		for j:=0; j<len(reviewGroups[i].Students); j++ {
			student, _ := services.GetStudent(ctx, reviewGroups[i].Students[j])
			if len(student.ReviewerGrades)==0 {
				noReviews++
			} else if len(student.ReviewerGrades) == len(reviewGroups[i].Reviewers) {
				allReviews++
			} else {
				someReviews++
			}
		}
	}

	type Response struct {
		NoReviews int		`json:"noReviews"`
		SomeReviews int		`json:"someReviews"`
		AllReviews int		`json:"allReviews"`
	}
	response := Response{NoReviews: noReviews, SomeReviews: someReviews, AllReviews: allReviews}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetReviewGroupForReviewer(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	user := context.Get(r, "user")

	// log.Println(user);

	mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	reviewerKey, err := datastore.DecodeKey(mapClaims["key"].(string))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	reviewGroup, reviewGroupKey, err := services.GetReviewGroupForReviewer(ctx, reviewerKey)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	var studentsArr = make([]models.Student, len(reviewGroup.Students))
	var studentGrades = make([]int, len(reviewGroup.Students))

	for i := 0; i < len(reviewGroup.Students); i++ {
    	student, _ := services.GetStudent(ctx, reviewGroup.Students[i])
    	studentsArr[i] = student

		grade, err := services.GetReviewerGradeForStudent(ctx, reviewerKey, reviewGroup.Students[i])
		if err != nil {
			studentGrades[i] = -1
		} else {
			studentGrades[i] = grade
		}
	}




	type Response struct {
		Key *datastore.Key		`json:"keys"`
		ReviewGroup models.ReviewGroup		`json:"users"`
		Students []models.Student 		`json:"students"`
		StudentGrades []int 		`json:"studentGrades"`
	}
	response := Response{Key: reviewGroupKey, ReviewGroup: reviewGroup, Students: studentsArr, StudentGrades: studentGrades}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetReviewerGradeForStudent(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	user := context.Get(r, "user")

	mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	reviewerKey, err := datastore.DecodeKey(mapClaims["key"].(string))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	studentKey, err := datastore.DecodeKey(mux.Vars(r)["studentKey"])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	grade, err := services.GetReviewerGradeForStudent(ctx, reviewerKey, studentKey)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	type Response struct {
		Grade int		`json:"grade"`
	}
	response := Response{Grade: grade}
	w.Header().Add("Access-Control-Allow-Origin", "*")	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func PostReviewerGradeForStudent(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	user := context.Get(r, "user")
	mapClaims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	reviewerKey, err := datastore.DecodeKey(mapClaims["key"].(string))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	studentKey, err := datastore.DecodeKey(dat["studentKey"].(string))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	reviewerGrade := int(dat["reviewerGrade"].(float64))

	err = services.UpdateReviewerGradeForStudent(ctx, reviewerKey, studentKey, reviewerGrade)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}