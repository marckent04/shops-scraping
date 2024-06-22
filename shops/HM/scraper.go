package HM

import (
	"shops-scraping/common"
)

type Scraper struct {
	url string
}

func (s Scraper) GetByKeywords(keywords string) (err error, articles []common.Article) {
	err, articles = getProducts(keywords)
	if err != nil {
		return err, nil
	}

	return
}

func NewScrapper() common.Scraper {
	return &Scraper{url: searchUrl}
}
