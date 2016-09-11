package main

import (
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/bmizerany/pat"
	"github.com/codegangsta/negroni"
	"github.com/spf13/viper"
)

var Settings map[string]*viper.Viper
var environments string

func newServer() *http.Server {
	n := negroni.New()

	// Middlewares
	if Settings["environments"].GetBool("log") {
		n.Use(negroni.NewLogger())
	}

	//n.UseFunc(recovery())
	n.Use(negroni.NewStatic(http.Dir("./public")))

	// Setup routes
	router := pat.New()
	//router.Get("/api/v1/jsonInfo", http.HandlerFunc(sampleHandler))

	// Add alive endpoint
	// router.Get("/alive", http.HandlerFunc(alive))
	// Add the router action
	n.UseHandler(router)
	Server := &http.Server{
		Addr:           ":" + Settings["environments"].GetString("server.port"),
		Handler:        n,
		MaxHeaderBytes: 1 << 20,
	}
	return Server
}

func InitConfigs() {
	environments = os.Getenv("ENV")
	if environments == "" {
		environments = "development"
	}
	Settings = make(map[string]*viper.Viper)
	Settings["environments"] = viper.New()
	Settings["environments"].SetConfigType("yml")
	Settings["environments"].AddConfigPath("config/environments")
	Settings["environments"].SetConfigName(environments)
	err := Settings["environments"].ReadInConfig()
	if err != nil {
		fmt.Errorf("Fatal error config file: %s \n", err)
	}
}

func main() {
	InitConfigs()
	InitApp()
	Server := newServer()
	log.Printf("Start serving on %s", Server.Addr)
	log.Println(Server.ListenAndServe())
}

func InitApp() {
	// add here plugins and middleware
}
