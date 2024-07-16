package PULLBEAR

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
	log.Printf("%s products getting in progress ...", shopName)

	page := browser.MustPage(fmt.Sprintf(searchUrl, genders[p.Gender], p.Keywords)).MustWaitDOMStable()

	acceptCookiesBtn := "#onetrust-accept-btn-handler"
	err = page.WaitElementsMoreThan(acceptCookiesBtn, 0)
	if err != nil {
		return err, nil
	}

	page.MustElement(acceptCookiesBtn).MustClick()

	err = page.Mouse.Scroll(0, 6000, 10)
	if err != nil {
		log.Error("error when scroll", err)
	}

	grid := getSearchGrid(page)

	if grid.MustHas(".results") {
		foundArticles := getArticlesSD(grid)
		for _, node := range foundArticles {
			articles = append(articles, rodeToArticle(node))
		}
	}

	log.Printf("%s products getting finished ...", shopName)

	return
}

func NewScrapper() common.Scraper {
	return &Scraper{url: searchUrl}
}
