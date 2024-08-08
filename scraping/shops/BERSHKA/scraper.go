package BERSHKA

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"shops-scraping/scraping/Browser"
	"shops-scraping/scraping/common"
	"shops-scraping/shared"
	"strings"
)

type Scraper struct {
	url string
}

func (s Scraper) GetByKeywords(p common.SearchParams) (err error, articles []shared.Article) {
	log.Printf("%s products getting in progress ...", shopName)

	browser := Browser.GetInstance()

	page := browser.MustPage(fmt.Sprintf("%s/%s?gender=%s", searchUrl, p.Keywords, genders[p.Gender])).MustWaitDOMStable()
	defer page.MustClose()

	isEmpty := page.MustHas(".search-empty__empty")

	if isEmpty {
		return
	}

	foundArticles := page.MustElements(articleSelector)

	for _, node := range foundArticles {
		h, _ := node.HTML()

		isArticle := strings.Contains(h, "data-qa-anchor=\"productItemText\"")

		if !isArticle {
			continue
		}

		articles = append(articles, rodeToArticle(node))
	}

	log.Printf("%s articles getting finished", shopName)

	if err != nil {
		log.Println(err.Error())
	}

	return
}

func NewScrapper() common.Scraper {
	return &Scraper{url: searchUrl}
}
