package main

import (
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"shops-scraping/scraping/Browser"
	"shops-scraping/webservice"
	articlesController "shops-scraping/webservice/articles"
)

func main() {

	setupEnv()

	Browser.CreateInstance()
	defer Browser.GetInstance().MustClose()

	startWebserver()

}

func setupEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func startWebserver() {
	webservice.ServeFrontend()
	webservice.RegisterApiRoutes(articlesController.Routes)

	port := os.Getenv("PORT")

	log.Println("Welcome to shop scraper api")
	log.Println("server is launching on port ", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
