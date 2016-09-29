package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2/google"
	storage "google.golang.org/api/storage/v1"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"

	"Xtern-Matching/models"
)


//8 MB file limit
const MAX_MEMORY = 8 * 1024 * 1024

func PostPDF(w http.ResponseWriter,r *http.Request) {


	bucketName := "xtern-matching.appspot.com"
	projectID := "xtern-matching"
	//bucketName := "xtern-matching-143216.appspot.com"//DEV Server
	//projectID := "xtern-matching-143216"
	id,_ := mux.Vars(r)["id"]

	//Make sure pdf is less than 8 MB
	if err := r.ParseMultipartForm(MAX_MEMORY); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	//Fetch file from formdata
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()

	//Get context and storage service
	ctx := appengine.NewContext(r)
	client, err := google.DefaultClient(ctx, storage.DevstorageFullControlScope)
	if err != nil {
			http.Error(w, err.Error(), 500)
			return
	}
	service, err := storage.New(client)
	if err != nil {
			http.Error(w, err.Error(), 500)
			return
	}

	//Access Bucket and see if it exists
	if _, err := service.Buckets.Get(bucketName).Do(); err == nil {
		log.Printf("Bucket %s already exists - skipping buckets.insert call.", bucketName)
	} else {
			// Create a bucket.
			if res, err := service.Buckets.Insert(projectID, &storage.Bucket{Name: bucketName}).Do(); err == nil {
					log.Printf("Created bucket %v at location %v\n\n", res.Name, res.SelfLink)
			} else {
				http.Error(w, err.Error(), 500)
				return
			}
	}

	//Delete old resume copy if it exists
	if err := service.Objects.Delete(bucketName, id + ".pdf").Do(); err != nil {
			// If the object exists but wasn't deleted, the bucket deletion will also fail.
			log.Printf("Could not delete object during cleanup: %v\n\n", err)
	} else {
			log.Printf("Successfully deleted %s/%s during cleanup.\n\n", bucketName, id)
	}

	//Insert new resume copy
	object := &storage.Object{Name: id + ".pdf"}
	res, err := service.Objects.Insert(bucketName, object).Media(file).Do()
	if err == nil {
			log.Printf("Created object %v at location %v\n\n", res.Name, res.SelfLink)
	} else {
			log.Printf("Objects.Insert failed: %v", err)
	}

	//Update student record to include resume link
	num_id, _ := strconv.ParseInt(id,10,64)
	studentKey := datastore.NewKey(ctx, "Student", "", num_id, nil)
	var student models.Student
	if err := datastore.Get(ctx, studentKey, &student); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	student.Resume = res.MediaLink
	if _, err := datastore.Put(ctx, studentKey, &student); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
