package HM

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"shops-scraping/scraping/Browser"
	"shops-scraping/scraping/common"
	"shops-scraping/shared"
)

type Scraper struct {
	url string
}

func (s Scraper) GetByKeywords(p common.SearchParams) (err error, articles []shared.Article) {
	msg := fmt.Sprintf("Search for %s at %s for %s", p.Keywords, shopName, p.Gender)
	log.Println(msg, " started")

	browser := Browser.GetInstance()

	page := browser.MustPage(fmt.Sprintf("%s?q=%s&department=%s", searchUrl, p.Keywords, genders[p.Gender])).MustWaitDOMStable()
	defer page.MustClose()

	if !page.MustHas(productsListSelector) {
		return
	}

	page = page.MustWaitElementsMoreThan(productSelector, 5)

	page.Mouse.Scroll(10, 10000, 10)

	foundArticles := page.MustElements(productSelector)

	log.Println(msg, " finished")

	for _, v := range foundArticles {
		articles = append(articles, rodeToArticle(v))
	}
	return
}

func NewScrapper() common.Scraper {
	return &Scraper{url: searchUrl}
}
