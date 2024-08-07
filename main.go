package main

import (
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"shops-scraping/database"
	"shops-scraping/scraping/Browser"
	"shops-scraping/webservice"
	articlesController "shops-scraping/webservice/articles"
	cartController "shops-scraping/webservice/cart"
)

func main() {

	setupEnv()

	Browser.CreateInstance()
	defer Browser.GetInstance().MustClose()

	database.Connect("test.db")

	startWebserver()

}

func setupEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var errorLogs string

	dbVariables := []string{
		"PORT",
	}

	for _, variable := range dbVariables {
		if os.Getenv(variable) == "" {
			errorLogs = fmt.Sprintf("%s\n%s", errorLogs, fmt.Sprintf("%s MUST BE DEFINED", variable))
		}
	}

	if errorLogs != "" {
		log.Fatal(errorLogs)
	}
}

func startWebserver() {
	webservice.ServeFrontend()
	webservice.RegisterApiRoutes(articlesController.Routes)
	webservice.RegisterApiRoutes(cartController.Routes)

	log.Println("Welcome to shop scraper app")

	port := os.Getenv("PORT")
	log.Println("server is launching on port ", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
