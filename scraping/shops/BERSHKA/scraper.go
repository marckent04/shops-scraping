package BERSHKA

import (
	"fmt"
	"github.com/go-rod/rod"
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

	page := browser.MustPage(fmt.Sprintf("%s/%s?gender=%s", searchUrl, p.Keywords, genders[p.Gender]))
	defer page.MustClose()
	defer log.Printf("%s articles getting finished", shopName)

	hasResults := s.waitForLoad(page)
	if !hasResults {
		return
	}

	foundArticles := page.MustElements(articleSelector)
	for _, node := range foundArticles {
		article, err := rodeToArticle(node)
		if err == nil {
			articles = append(articles, article)
		}
	}

	if err != nil {
		log.Println(err.Error())
	}

	return
}

func (s Scraper) waitForLoad(page *rod.Page) (hasResults bool) {
	var ch = make(chan bool, 1)

	go func() {
		err := page.WaitElementsMoreThan(articleSelector, 0)
		if err == nil {
			ch <- true
		}
	}()
	go func() {
		err := page.WaitElementsMoreThan(".search-empty__empty", 0)
		if err == nil {
			ch <- false
		}
	}()

	hasResults = <-ch

	return
}

func NewScrapper() common.Scraper {
	return &Scraper{url: searchUrl}
}
