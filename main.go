package main

import (
	"encoding/json"
	"log"
	HM "shop-scraping/shops/HM"
)

func main() {
	HMScraper := HM.NewScrapper()

	err, articles := HMScraper.GetByKeywords("polo")
	if err != nil {
		log.Fatal(err)
	}

	//if len(articles) != 36 {
	//	log.Printf("Articles got %d", len(articles))
	//	log.Panic("all articles must be get")
	//}
	//
	//for _, article := range articles {
	//	fmt.Printf("%s = %s\n", article.Name, article.Price)
	//}

	_, err = json.Marshal(articles)
	if err != nil {
		log.Fatal(err)
	}

	//log.Println(string(data))
}
