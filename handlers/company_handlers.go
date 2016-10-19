package handlers

import (
	"net/http"
	"google.golang.org/appengine"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"Xtern-Matching/models"
	"Xtern-Matching/handlers/services"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	// "github.com/gorilla/context"
	// "github.com/dgrijalva/jwt-go"
)

func AddStudent(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)
	// claims := context.Get(r, "user").(*jwt.Token).Claims.(jwt.MapClaims)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	studentId :=  int64(dat["studentId"].(float64));
	companyId :=  int64(dat["id"].(float64));

	// if(claims["org"] == companyId) {

		_, err := services.AddStudentIdToCompanyList(ctx, companyId, studentId)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(http.StatusOK)
	// } else {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// }
}

func SwitchStudents(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)
	// claims := context.Get(r, "user").(*jwt.Token).Claims.(jwt.MapClaims)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	student1Id :=  int64(dat["student1Id"].(float64));
	student2Id :=  int64(dat["student2Id"].(float64));
	companyId :=  int64(dat["id"].(float64));

	// if(claims["org"] == companyId) {
		_, err := services.SwitchStudentIdsInCompanyList(ctx, companyId, student1Id, student2Id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(http.StatusOK)
	// } else {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// }	
}

func RemoveStudent(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	studentId :=  int64(dat["studentId"].(float64));
	companyId :=  int64(dat["id"].(float64));

	_, err := services.RemoveStudentIdFromCompanyList(ctx, companyId, studentId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func PostCompany(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	var companies []models.Company
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&companies); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	for _, company := range companies {
		_, err := services.NewCompany(ctx, company)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}

func GetCompany(w http.ResponseWriter,r *http.Request) {
	ctx := appengine.NewContext(r)

	if id, ok := mux.Vars(r)["Id"]; ok {
		num_id, _ := strconv.ParseInt(id, 10, 64)
		company, err := services.GetCompany(ctx, num_id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(company)
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func GetCurrentCompany(w http.ResponseWriter,r *http.Request) {
	// ctx := appengine.NewContext(r)

	var dat map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tokenString := dat["token"].(string)
	// // var token []byte
	// // token, _ = jwt.DecodeSegment(tokenString)

	// // Parse the token
	// token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsExample{}, func(token *jwt.Token) (interface{}, error) {
 //    // since we only use the one private key to sign the tokens,
 //    // we also only use its public counter part to verify
	// 	return verifyKey, nil
	// })
	// // fatal(err)

	// claims := token.Claims.(*CustomClaimsExample)
	// // fmt.Println(claims.CustomerInfo.Name)


	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    // Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("My Secret"), nil
		})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}




	// func DecodeSegment(seg string) ([]byte, error)

	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) ([]byte, error) {

 //        // fmt.Printf("parsed token obj: %v\n", token)
 //        // publicKey, err := ioutil.ReadFile("keys/app.rsa.pub")
 //        // if err != nil {
 //        //     return nil, fmt.Errorf("Error reading public key")
 //        // }
 //        var publicKey = ([]byte("My Secret"));

 //        return publicKey, nil
 //    })




	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
	w.WriteHeader(http.StatusInternalServerError)


	// token :=  int64(dat["studentId"].(float64));

	// _, err := services.RemoveStudentIdFromCompanyList(ctx, companyId, studentId)
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
}