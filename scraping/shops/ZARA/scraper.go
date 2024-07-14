package ZARA

import (
	"github.com/go-rod/rod"
	"shops-scraping/scraping/common"
	"shops-scraping/shared"
)

type Scraper struct {
	url string
}

func (s Scraper) GetByKeywords(browser *rod.Browser, keyword string) (err error, articles []shared.Article) {
	err, articles = getProducts(browser, keyword)
	if err != nil {
		return err, nil
	}
	return
}

func NewScrapper() common.Scraper {
	return &Scraper{url: searchUrl}
}
