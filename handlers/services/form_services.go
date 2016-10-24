package services

//import (
//	"net/http"
//	"time"
//	"github.com/parnurzeal/gorequest"
//	"fmt"
//)
//
//func Authorize(client_id string, redirect_uri string, code string) {
//	res, body, err := gorequest.New().
//		Get("/api/v2/oauth2/authorize?" +
//			"client_id=" + client_id +
//			"redirect_uri=" + redirect_uri +
//			"response_type=" + code).
//		Set("Host","www.formstack.com").
//		End()
//	if err != nil {
//		fmt.Println(err)
//	}
//}
//
//func GetFormDetails(formId string)  {
//	res, body, err := gorequest.New().
//		Get("http://www.formstack.com//api/v2/form/"+formId+".json").
//		Set("Authorization","Bearer "+formId).
//		End()
//	if err != nil {
//		fmt.Println(err)
//	}
//}
