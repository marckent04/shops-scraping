package articlesController

import (
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"net/http"
	"shops-scraping/scraping/common"
	"shops-scraping/scraping/shops/BERSHKA"
	"shops-scraping/scraping/shops/HM"
	"shops-scraping/scraping/shops/SHEIN"
	"shops-scraping/shared"
)

func searchByKeywords(rsp http.ResponseWriter, req *http.Request) {
	var allArticles []shared.Article
	keyword := req.URL.Query().Get("keyword")

	totalShops := 3

	articlesCh := make(chan []shared.Article, totalShops)

	go getArticlesByKeywords(articlesCh, SHEIN.NewScrapper(), keyword)
	go getArticlesByKeywords(articlesCh, HM.NewScrapper(), keyword)
	go getArticlesByKeywords(articlesCh, BERSHKA.NewScrapper(), keyword)

	for i := 1; i <= totalShops; i++ {
		allArticles = append(allArticles, <-articlesCh...)
	}
	close(articlesCh)

	indent, err := json.MarshalIndent(allArticles, "", " ")
	if err != nil {
		rsp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rsp, "Error during response formatage")
		return
	}

	rsp.Header().Set("Content-Type", "application/json")
	rsp.Header().Set("Accept-Charset", "UTF-8")
	rsp.WriteHeader(http.StatusOK)
	log.Printf("response done for search %s", keyword)
	rsp.Write(indent)
}

func getSheinArticles(ch chan []shared.Article, keyword string) {
	scper := SHEIN.NewScrapper()

	err, art := scper.GetByKeywords(keyword)
	if err != nil {
		log.Println(err)
		return
	}

	ch <- art
}

func getHMArticles(ch chan<- []shared.Article, keyword string) {
	scper := HM.NewScrapper()

	err, art := scper.GetByKeywords(keyword)
	if err != nil {
		return
	}

	ch <- art
}

func getArticlesByKeywords(ch chan<- []shared.Article, scraper common.Scraper, keyword string) {

	err, art := scraper.GetByKeywords(keyword)
	if err != nil {
		return
	}

	ch <- art
}
