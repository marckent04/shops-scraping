package ZARA

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

	browser := Browser.GetInstance()

	gender := genders[p.Gender]

	log.Printf("%s products for %s getting in progress ...", shopName, gender)

	page := browser.MustPage(fmt.Sprintf(searchUrl, p.Keywords, gender)).MustWaitDOMStable()
	defer page.MustClose()

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
