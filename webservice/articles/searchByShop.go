package articlesController

import (
	"encoding/json"
	"net/http"
	"shops-scraping/scraping/shops/BERSHKA"
	"shops-scraping/scraping/shops/HM"
	"shops-scraping/scraping/shops/PULLBEAR"
	"shops-scraping/scraping/shops/ZARA"
	"shops-scraping/shared"
	"slices"
	"strings"
)

var supportedShops = []shared.Shop{
	shared.HM,
	shared.BERSHKA,
	shared.ZARA,
	shared.PULLANDBEAR,
}

func searchByShops(rsp http.ResponseWriter, req *http.Request) {
	keyword := req.URL.Query().Get("q")
	shopsQuery := req.URL.Query().Get("shops")

	var articles []shared.Article

	shops := getShopsFromQuery(shopsQuery)

	if len(shops) == 0 {
		rsp.WriteHeader(http.StatusBadRequest)
		rsp.Write([]byte("NO SUPPORTED SHOPS FOUND"))
		return
	}

	articlesChan := make(chan []shared.Article, len(shops))
	defer close(articlesChan)

	for _, shop := range shops {
		fetchArticles(shop, keyword, articlesChan)
	}

	for i := 0; i < len(shops); i++ {
		articles = append(articles, <-articlesChan...)
	}

	indent, _ := json.MarshalIndent(articles, "", " ")
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

func fetchArticles(shop shared.Shop, keywords string, articlesChan chan<- []shared.Article) {
	switch shop {
	/*case shared.SHEIN:
	getArticlesByKeywords(articlesChan, SHEIN.NewScrapper(), keyword)
	break*/
	case shared.HM:
		getArticlesByKeywords(articlesChan, HM.NewScrapper(), keywords)
		break
	case shared.BERSHKA:
		getArticlesByKeywords(articlesChan, BERSHKA.NewScrapper(), keywords)
		break
	case shared.ZARA:
		getArticlesByKeywords(articlesChan, ZARA.NewScrapper(), keywords)
		break
	case shared.PULLANDBEAR:
		getArticlesByKeywords(articlesChan, PULLBEAR.NewScrapper(), keywords)
		break
	}
}
