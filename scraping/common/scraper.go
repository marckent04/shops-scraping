package common

import "shops-scraping/shared"

type Scraper interface {
	GetByKeywords(keywords string) (error, []shared.Article)
}
