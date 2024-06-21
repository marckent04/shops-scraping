package scraping

import "shops-scraping/article"

type Scraper interface {
	GetByKeywords(keywords string) (error, []article.Article)
}
