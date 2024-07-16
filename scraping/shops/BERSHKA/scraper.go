package BERSHKA

import (
	"fmt"
	"github.com/go-rod/rod"
	log "github.com/sirupsen/logrus"
	"shops-scraping/scraping/common"
	"shops-scraping/shared"
	"strings"
)

type Scraper struct {
	url string
}

func (s Scraper) GetByKeywords(browser *rod.Browser, p common.SearchParams) (err error, articles []shared.Article) {
	log.Printf("%s products getting in progress ...", shopName)

	page := browser.MustPage(fmt.Sprintf("%s/%s?gender=%s", searchUrl, p.Keywords, genders[p.Gender])).MustWaitDOMStable()

	isEmpty := page.MustHas(".search-empty__empty")

	if isEmpty {
		return
	}

	page = page.MustWaitElementsMoreThan(articleSelector, 4)

	page.Mouse.Scroll(10, 10000, 15)

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
