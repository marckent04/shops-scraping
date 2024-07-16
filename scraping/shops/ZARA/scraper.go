package ZARA

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

	gender := genders[p.Gender]

	log.Printf("%s products for %s getting in progress ...", shopName, gender)

	page := browser.MustPage(fmt.Sprintf(searchUrl, p.Keywords, gender)).MustWaitDOMStable()

	if page.MustHas(".zds-empty-state__title") {
		return
	}

	page = page.MustWaitElementsMoreThan(articleSelector, 1)

	err = page.Mouse.Scroll(10, 10000, 15)
	if err != nil {
		return err, nil
	}

	foundArticles := page.MustElements(articleSelector)

	for _, node := range foundArticles {
		if !node.MustHas(".money-amount__main") {
			continue
		}

		articles = append(articles, rodeToArticle(node))
	}
	return
}

func NewScrapper() common.Scraper {
	return &Scraper{url: searchUrl}
}
