package main

import (
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"shops-scraping/database"
	"shops-scraping/scraping/Browser"
	"shops-scraping/webservice"
	articles "shops-scraping/webservice/articles"
	"shops-scraping/webservice/cart"
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
	log.Println("Welcome to shop scraper app")
	webservice.ServeFrontend()

	port := os.Getenv("PORT")
	app := webservice.NewHttpRouter()

	app.SetGlobalPrefix("/api")
	app.Get("/articles", articles.SearchByShops)
	app.Get("/cart", cart.GetCart)
	app.Delete("/cart/line", cart.DeleteCartLine)
	app.Delete("/cart", cart.ClearCart)
	app.Post("/cart", cart.AddToCart)

	app.Listen(port)

}
