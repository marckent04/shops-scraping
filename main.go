package main

import (
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"shops-scraping/database"
	"shops-scraping/scraping/Browser"
	"shops-scraping/webservice"
)

func main() {

	setupEnv()

	Browser.CreateInstance()
	defer Browser.GetInstance().MustClose()

	database.Connect(os.Getenv("DB_PATH"))

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
	log.Println("Welcome to shop scraper app")
	webservice.Start()
}
