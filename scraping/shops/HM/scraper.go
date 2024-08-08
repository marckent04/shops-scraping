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

	page := browser.MustPage(fmt.Sprintf("%s?q=%s&department=%s", searchUrl, p.Keywords, genders[p.Gender]))
	defer page.MustClose()

	err = page.WaitElementsMoreThan(productSelector, 5)
	if err != nil {
		return
	}

	if !page.MustHas(productsListSelector) {
		return
	}
	common.CloseCookieDialog(page)

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
