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

	common.CloseCookieDialog(page)

	// TODO: current time: 2.3s [a optimiser]
	foundArticles := page.MustElements(articleSelector)
	for _, node := range foundArticles {
		if !node.MustHas(".money-amount__main") {
			continue
		}

		if node.MustHas(".products-category-grid-media-carousel-placeholder__loader") {
			node.MustHover()
		}

		articles = append(articles, rodeToArticle(node))
	}
	return
}

func NewScrapper() common.Scraper {
	return &Scraper{url: searchUrl}
}
