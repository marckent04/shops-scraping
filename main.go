package main

import (
	"encoding/json"
	"log"
	HM "shops-scraping/shops/HM"
)

func main() {
	HMScraper := HM.NewScrapper()

	err, articles := HMScraper.GetByKeywords("polo")
	if err != nil {
		log.Fatal(err)
	}

	articlesJson, err := json.Marshal(articles)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(articlesJson))
}
