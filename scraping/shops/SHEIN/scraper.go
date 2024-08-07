package SHEIN

import (
	"shops-scraping/shared"
)

type Scraper struct {
	url string
}

func (s Scraper) GetByKeywords(keywords string) (err error, articles []shared.Article) {
	articles = getProducts(keywords)

	if err != nil {
		return err, nil
	}

	return
}
