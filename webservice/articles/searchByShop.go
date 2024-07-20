package articlesController

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"shops-scraping/scraping/common"
	"shops-scraping/scraping/shops/BERSHKA"
	"shops-scraping/scraping/shops/HM"
	"shops-scraping/scraping/shops/PULLBEAR"
	"shops-scraping/scraping/shops/ZARA"
	"shops-scraping/shared"
	"slices"
	"strings"
	"time"
)

var supportedShops = []shared.Shop{
	shared.HM,
	shared.BERSHKA,
	shared.ZARA,
	shared.PULLANDBEAR,
}

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

	shops := getShopsFromQuery(shopsQuery)

	if len(shops) == 0 {
		rsp.WriteHeader(http.StatusBadRequest)
		rsp.Write([]byte("NO SUPPORTED SHOPS FOUND"))
		return
	}

	articlesChan := make(chan []shared.Article, len(shops))
	defer close(articlesChan)

	params := common.NewSearchParams(gender, keyword)

	for _, shop := range shops {
		go fetchArticles(shop, params, articlesChan)
	}

	for i := 0; i < len(shops); i++ {
		articles = append(articles, <-articlesChan...)
	}

	indent, _ := json.MarshalIndent(articles, "", " ")
	elapsed := time.Since(start)

	log.Println("time elapsed : ", elapsed)
	rsp.Write(indent)
}

func getShopsFromQuery(query string) (shops []shared.Shop) {
	rShops := strings.Split(query, ",")
	for _, shop := range supportedShops {
		if slices.Contains(rShops, shop) {
			shops = append(shops, shop)
		}
	}
	return
}

func fetchArticles(shop shared.Shop, params common.SearchParams, ch chan<- []shared.Article) {
	scraper := getScraper(shop)
	if scraper == nil {
		return
	}

	err, arts := scraper.GetByKeywords(params)

	if err != nil {
		return
	}

	ch <- arts
}

func getScraper(shop shared.Shop) common.Scraper {
	switch shop {
	case shared.HM:
		return HM.NewScrapper()
	case shared.BERSHKA:
		return BERSHKA.NewScrapper()
	case shared.ZARA:
		return ZARA.NewScrapper()
	case shared.PULLANDBEAR:
		return PULLBEAR.NewScrapper()
	}
	return nil
}
