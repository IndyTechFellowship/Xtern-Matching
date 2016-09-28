package handlers

import (
	"log"
	"os"
	"strconv"
	"net/http"
	"io/ioutil"
	
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	storage "google.golang.org/api/storage/v1"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"Xtern-Matching/models"
)

/*
func GetPdf(w http.ResponseWriter,r *http.Request) {
	id,_ := mux.Vars(r)["id"]
	if res, err := service.Objects.Get(*bucketName, id).Do(); err == nil {
			fmt.Printf("The media download link for %v/%v is %v.\n\n", *bucketName, res.Name, res.MediaLink)
	} else {
			fatalf(service, "Failed to get %s/%s: %s.", *bucketName, objectName, err)
	}
	if !restoreOriginalState(service) {
			//os.Exit(1)
	}
}
*/

//8 MB file limit
const MAX_MEMORY = 8 * 1024 * 1024

func PostPDF(w http.ResponseWriter,r *http.Request) {

	bucketName := "Resumes"
	projectID := "xtern-matching-143216"
	
	id,_ := mux.Vars(r)["id"]
	log.Println(id)
	
	//Make sure pdf is less than 8 MB
	if err := r.ParseMultipartForm(MAX_MEMORY); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	
	//Get file
	upfile,_ := ioutil.TempFile(os.TempDir(), id)
	defer os.Remove(upfile.Name())
	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, _ := fileHeader.Open()
			buf, _ := ioutil.ReadAll(file)
			upfile.Write(buf)
		}
	}
	
	//Access Bucket and see if it exists
	client, err := google.DefaultClient(context.Background(), storage.DevstorageFullControlScope)
	if err != nil {
			log.Fatalf("Unable to get default client: %v", err)
			return
	}
	service, err := storage.New(client)
	if err != nil {
			log.Fatalf("Unable to create storage service: %v", err)
			return
	}
	if _, err := service.Buckets.Get(bucketName).Do(); err == nil {
		log.Printf("Bucket %s already exists - skipping buckets.insert call.", bucketName)
	} else {
			// Create a bucket.
			if res, err := service.Buckets.Insert(projectID, &storage.Bucket{Name: bucketName}).Do(); err == nil {
					log.Printf("Created bucket %v at location %v\n\n", res.Name, res.SelfLink)
			} else {
					log.Fatalf("Failed creating bucket %s: %v", bucketName, err)
			}
	}
	
	//Delete old resume copy if it exists
	if err := service.Objects.Delete(bucketName, id).Do(); err != nil {
			// If the object exists but wasn't deleted, the bucket deletion will also fail.
			log.Printf("Could not delete object during cleanup: %v\n\n", err)
	} else {
			log.Printf("Successfully deleted %s/%s during cleanup.\n\n", bucketName, id)
	}	
	
	//Insert new resume copy
	object := &storage.Object{Name: id}
	res, err := service.Objects.Insert(bucketName, object).Media(upfile).Do()
	if err == nil {
			log.Printf("Created object %v at location %v\n\n", res.Name, res.SelfLink)
	} else {
			log.Printf("Objects.Insert failed: %v", err)
	}
	
	//Update student record to include resume link
	ctx := appengine.NewContext(r)
	num_id, _ := strconv.ParseInt(id,10,64)
	studentKey := datastore.NewKey(ctx, "Student", "", num_id, nil)
	var student models.Student
	if err := datastore.Get(ctx, studentKey, &student); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	student.Resume = res.SelfLink
	if _, err := datastore.Put(ctx, studentKey, &student); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	
}