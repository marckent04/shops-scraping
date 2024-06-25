package articlesController

import (
	"encoding/json"
	"net/http"
	"shops-scraping/scraping/shops/BERSHKA"
	"shops-scraping/scraping/shops/HM"
	"shops-scraping/scraping/shops/SHEIN"
	"shops-scraping/shared"
)

func searchByShop(rsp http.ResponseWriter, req *http.Request) {
	keyword := req.URL.Query().Get("keyword")
	shop := req.PathValue("shop")

	articlesChan := make(chan []shared.Article, 1)

	switch shop {
	case shared.SHEIN:
		getArticlesByKeywords(articlesChan, SHEIN.NewScrapper(), keyword)
		break
	case shared.HM:
		getArticlesByKeywords(articlesChan, HM.NewScrapper(), keyword)
		break
	case shared.BERSHKA:
		getArticlesByKeywords(articlesChan, BERSHKA.NewScrapper(), keyword)
		break
	}

	indent, _ := json.MarshalIndent(<-articlesChan, "", " ")
	rsp.Write(indent)
}
