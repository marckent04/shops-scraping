package articlesController

import (
	"encoding/json"
	"github.com/go-rod/rod"
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

	var articles []shared.Article

	shops := getShopsFromQuery(shopsQuery)

	if len(shops) == 0 {
		rsp.WriteHeader(http.StatusBadRequest)
		rsp.Write([]byte("NO SUPPORTED SHOPS FOUND"))
		return
	}

	browser := rod.New().MustConnect()
	defer browser.MustClose()

	articlesChan := make(chan []shared.Article, len(shops))
	defer close(articlesChan)

	for _, shop := range shops {
		go fetchArticles(browser, shop, keyword, articlesChan)
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

func fetchArticles(browser *rod.Browser, shop shared.Shop, keywords string, articlesChan chan<- []shared.Article) {
	switch shop {
	/*case shared.SHEIN:
	getArticlesByKeywords(articlesChan, SHEIN.NewScrapper(), keyword)
	break*/
	case shared.HM:
		getArticlesByKeywords(browser, articlesChan, HM.NewScrapper(), keywords)
		break
	case shared.BERSHKA:
		getArticlesByKeywords(browser, articlesChan, BERSHKA.NewScrapper(), keywords)
		break
	case shared.ZARA:
		getArticlesByKeywords(browser, articlesChan, ZARA.NewScrapper(), keywords)
		break
	case shared.PULLANDBEAR:
		getArticlesByKeywords(browser, articlesChan, PULLBEAR.NewScrapper(), keywords)
		break
	}
}

func getArticlesByKeywords(browser *rod.Browser, ch chan<- []shared.Article, scraper common.Scraper, keyword string) {
	err, art := scraper.GetByKeywords(browser, keyword)
	if err != nil {
		return
	}

	ch <- art
}
