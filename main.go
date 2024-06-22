package main

import (
	"encoding/json"
	"log"
	"shops-scraping/shops/SHEIN"
)

func main() {
	HMScraper := SHEIN.NewScrapper()

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
