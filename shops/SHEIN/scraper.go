package SHEIN

import (
	"shops-scraping/article"
	"shops-scraping/scraping"
)

type Scraper struct {
	url string
}

func (s Scraper) GetByKeywords(keywords string) (err error, articles []article.Article) {
	err, articles = getProducts(keywords)
	if err != nil {
		return err, nil
	}

	return
}

func NewScrapper() scraping.Scraper {
	return &Scraper{url: searchUrl}
}
