package services

import (
	"Xtern-Matching/models"
	"archive/zip"
	"io"
	// "io/ioutil"
	"log"
	"net/http"

	"Xtern-Matching/handlers/services/csv"

	"os"
	"strconv"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/storage/v1"
	"google.golang.org/appengine/datastore"
	"bytes"
	// "google.golang.org/appengine/file"
	// "fmt"
	"google.golang.org/appengine/urlfetch"
	//"fmt"
)

func GetStudents(ctx context.Context, parent *datastore.Key) ([]models.Student, []*datastore.Key, error) {
	q := datastore.NewQuery("Student")
	if parent != nil {
		q = datastore.NewQuery("Student").Ancestor(parent)
	}
	var students []models.Student
	keys, err := q.GetAll(ctx, &students)
	if err != nil {
		return nil, nil, err
	}

	return students, keys, nil
}

/*
Exports all student resumes in the Database.
Queries all students and exports them
 */
func ExportAllResumes(ctx context.Context) (*bytes.Buffer, error) {
	students, _, err := GetStudents(ctx, nil)
	if err != nil {
		return nil, err
	}
	return ExportResumes(ctx, students)
}

/*
Exports a slice of students as archive.pdf.
Useful for testing service to minimize the number of pdf GET requests
 */
func ExportResumes(ctx context.Context, students []models.Student) (*bytes.Buffer, error) {

	client := urlfetch.Client(ctx)
	buf := new(bytes.Buffer)

	archive := zip.NewWriter(buf)
	defer archive.Close()
	for _, student := range students {
		// Get the resume and write it
		resp, err := client.Get(student.Resume)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		f, err := archive.Create(student.Email + ".pdf")
		io.Copy(f, resp.Body)
	}

	return buf, nil
}

func GetStudent(ctx context.Context, studentKey *datastore.Key) (models.Student, error) {
	//studentKey := datastore.NewKey(ctx, "Student", "", _id, nil)
	var student models.Student
	err := datastore.Get(ctx, studentKey, &student)
	if err != nil {
		return models.Student{}, err
	}
	return student, nil
}

func ExportStudents(ctx context.Context) ([]byte, error) {
	students, _, err := GetStudents(ctx, nil)
	if err != nil {
		return nil, err
	}
	output, err := csv.Marshal(students)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func NewStudent(ctx context.Context, student models.Student) (int, error) {

	key := datastore.NewIncompleteKey(ctx, "Student", nil)
	student.Active = true

	//TODO make this done in a single put
	key, err := datastore.Put(ctx, key, &student)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	file, err := os.Open("public/sample.pdf")
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer file.Close()

	/* resumeURL, err := addResume(ctx, key.IntID(), file)
	if err != nil {
		log.Println("Error uploading resume")
		return http.StatusInternalServerError, err
	} */
	student.Resume = "http://localhost:8080/public/sample.pdf"//resumeURL
	_, err = datastore.Put(ctx, key, &student)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}

func addResume(ctx context.Context, studentId int64, file io.Reader) (string, error) {
	var bucketName string
	var projectID string
	if os.Getenv("XTERN_ENVIRONMENT") != "production" {
		bucketName = "xtern-matching-143216.appspot.com" //DEV Server
		projectID = "xtern-matching-143216"
	} else {
		bucketName = "xtern-matching.appspot.com"
		projectID = "xtern-matching"
	}

	client, err := google.DefaultClient(ctx, storage.DevstorageFullControlScope)
	if err != nil {
		log.Println("Error getting storage client")
		return "", err
	}
	service, err := storage.New(client)
	if err != nil {
		log.Println("Error getting storage service")
		return "", err
	}

	//Access Bucket and see if it exists
	if _, err := service.Buckets.Get(bucketName).Do(); err == nil {
		log.Printf("Bucket %s already exists - skipping buckets.insert call.", bucketName)
	} else {
		// Create a bucket.
		if res, err := service.Buckets.Insert(projectID, &storage.Bucket{Name: bucketName}).Do(); err == nil {
			log.Printf("Created bucket %v at location %v\n\n", res.Name, res.SelfLink)
		} else {
			return "", err
		}
	}

	//Insert new resume copy
	object := &storage.Object{Name: strconv.FormatInt(studentId, 10) + ".pdf"}
	res, err := service.Objects.Insert(bucketName, object).Media(file).Do()
	if err == nil {
		log.Printf("Created object %v at location %v\n\n", res.Name, res.SelfLink)
	} else {
		log.Println("Error inserting into bucket")
		return "", err
	}

	return res.MediaLink, nil
}

/*
	Array Mapping
	firstName 		-> 0
	lastName  		-> 1
	email  			-> 2
	university  		-> 3
	major	  		-> 4
	gradYear  		-> 5
	workStatus  		-> 6
	gender  		-> 7
	skills  		-> 8
	githubUrl  		-> 9
	linkedinUrl  		-> 10
	personalWebiteUrl 	-> 11
	interestedIn  		-> 12
	resume			-> 13
	homeState		-> 14
	status  		-> 15
	active			-> 16
	grade			-> 17
*/
func AddMappedStudent(ctx context.Context, mapping []string, data map[string]interface{}) (int, error) {
	var student models.Student
	var err error

	for i := 0; i < 16; i++ {
		if i == 8 || i == 12 {
			continue
		} else if data[mapping[i]] == nil {
		data[mapping[i]] = ""
		}
	}
	student.FirstName = data[mapping[0]].(string)
	student.LastName = data[mapping[1]].(string)
	student.Email = data[mapping[2]].(string)
	student.University = data[mapping[3]].(string)
	student.Major = data[mapping[4]].(string)
	student.GradYear = data[mapping[5]].(string)
	student.WorkStatus = data[mapping[6]].(string)
	student.Gender = data[mapping[7]].(string)

	if data[mapping[8]] != nil {
		skills := data[mapping[8]].([]interface{})
		student.Skills = make([]models.Skill, len(skills))
		for i:=0; i < len(skills); i++ {
			student.Skills[i].Name = skills[i].(map[string]interface{})["name"].(string)
			student.Skills[i].Name = skills[i].(map[string]interface{})["category"].(string)
		}
	}

	student.Github = data[mapping[9]].(string)
	student.Linkin = data[mapping[10]].(string)
	student.PersonalSite = data[mapping[11]].(string)
	if data[mapping[12]] != nil {
		interests := data[mapping[12]].([]interface{})
		student.Interests = make([]string, len(interests))
		for i:=0; i < len(interests); i++ {
			student.Interests[i] = interests[i].(string)
		}
	}
	student.Resume = data[mapping[13]].(string)
	student.HomeState = data[mapping[14]].(string)
	student.Status = data[mapping[15]].(string)
	if data[mapping[16]] == nil {
		student.Active = true
	} else {
		student.Active, err = strconv.ParseBool(data[mapping[16]].(string))
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}
	if data[mapping[17]] == nil {
		student.Grade = 0
	} else {
		student.Grade = data[mapping[17]].(int)
	}

	return NewStudent(ctx, student)
}