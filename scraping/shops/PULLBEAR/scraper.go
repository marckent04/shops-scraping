package PULLBEAR

import (
	"shops-scraping/scraping/common"
	"shops-scraping/shared"
)

type Scraper struct {
	url string
}

func (s Scraper) GetByKeywords(keyword string) (err error, articles []shared.Article) {
	err, articles = getProducts(keyword)
	if err != nil {
		return err, nil
	}
	return
}

func NewScrapper() common.Scraper {
	return &Scraper{url: searchUrl}
}
