package PULLBEAR

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
	log.Printf("%s products getting in progress ...", shopName)

	browser := Browser.GetInstance()

	page := browser.MustPage(fmt.Sprintf(searchUrl, genders[p.Gender], p.Keywords)).MustWaitDOMStable()
	defer page.MustClose()

	common.CloseCookieDialog(page)

	grid := getSearchGrid(page)
	if grid.MustHas(".results") {
		// TODO: current duration: 929ms [to optimize]
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
