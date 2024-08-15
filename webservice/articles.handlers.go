package webservice

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"shops-scraping/scraping/common"
	"shops-scraping/scraping/shops"
	"shops-scraping/scraping/shops/BERSHKA"
	"shops-scraping/scraping/shops/HM"
	"shops-scraping/scraping/shops/PULLBEAR"
	"shops-scraping/scraping/shops/ZARA"
	"shops-scraping/shared"
	"slices"
	"strings"
	"time"
)

func searchByShops(rsp http.ResponseWriter, req *http.Request) {
	start := time.Now()

	keyword := req.URL.Query().Get("q")
	shopsQuery := req.URL.Query().Get("shops")
	gender := req.URL.Query().Get("gender")

	if gender == "" || keyword == "" {
		rsp.WriteHeader(http.StatusBadRequest)
		rsp.Write([]byte("keyword or gender missing"))
		return
	}

	var articles []shared.Article

	reqShops := getShopsFromQuery(shopsQuery)

	if len(reqShops) == 0 {
		rsp.WriteHeader(http.StatusBadRequest)
		rsp.Write([]byte("NO SUPPORTED SHOPS FOUND"))
		return
	}

	articlesChan := make(chan []shared.Article, len(reqShops))
	defer close(articlesChan)

	params := common.NewSearchParams(gender, keyword)

	for _, shop := range reqShops {
		go fetchArticles(shop, params, articlesChan)
	}

	for i := 0; i < len(reqShops); i++ {
		articles = append(articles, <-articlesChan...)
	}

	serveJsonResponse(rsp, articles)
	log.Println("time elapsed : ", time.Since(start))

}

func getShopsFromQuery(query string) (foundShops []shared.Shop) {
	rShops := strings.Split(query, ",")
	enabledShops := shops.GetEnabledShops()
	for _, shop := range rShops {
		if slices.ContainsFunc(enabledShops, func(enabledShop shops.Shop) bool {
			return enabledShop.Code == shop
		}) {
			foundShops = append(foundShops, shop)
		}
	}
	return
}

func fetchArticles(shop shared.Shop, params common.SearchParams, ch chan<- []shared.Article) {
	scraper := getScraper(shop)
	if scraper == nil {
		ch <- make([]shared.Article, 0)
		return
	}

	arts, err := scraper.GetByKeywords(params)
	if err != nil {
		ch <- make([]shared.Article, 0)
		return
	}

	ch <- arts
}

func getScraper(shop shared.Shop) common.Scraper {
	switch shop {
	case shared.HM:
		return HM.GetScrapper()
	case shared.BERSHKA:
		return BERSHKA.GetScrapper()
	case shared.ZARA:
		return ZARA.GetScrapper()
	case shared.PULLANDBEAR:
		return PULLBEAR.GetScrapper()
	}
	return nil
}
