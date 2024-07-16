package HM

import (
	"fmt"
	"github.com/go-rod/rod"
	log "github.com/sirupsen/logrus"
	"shops-scraping/scraping/common"
	"shops-scraping/shared"
)

type Scraper struct {
	url string
}

func (s Scraper) GetByKeywords(browser *rod.Browser, p common.SearchParams) (err error, articles []shared.Article) {
	msg := fmt.Sprintf("Search for %s at %s for %s", p.Keywords, shopName, p.Gender)
	log.Println(msg, " started")

	page := browser.MustPage(fmt.Sprintf("%s?q=%s&department=%s", searchUrl, p.Keywords, genders[p.Gender])).MustWaitDOMStable()

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
