package routes

import(
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"Xtern-Matching/models"
	//"fmt"
	"net/http")

//GET user/{id}
func GetUserById(w http.ResponseWriter, r *http.Request){
	
	
}

//POST user
func CreateUser(w http.ResponseWriter, r *http.Request){
	context := appengine.NewContext(r)
	k := datastore.NewKey(context, "Student", "stringID", 0, nil)
	student := new(models.Student)
	
	if err := datastore.Get(context, k, student); err != nil {
        http.Error(w, datastore.IntID(k), 500)
        return
    }
	
	student.FirstName = "Bobby"
	
	    if _, err := datastore.Put(context, k, student); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
	
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    //fmt.Fprintf(w, "old=%q\nnew=%q\n", old, e.Value)
}

func UserRoute(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		case "POST":
			CreateUser(w, r)
		default:
	}
}