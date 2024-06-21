package scraping

import "shop-scraping/article"

type Scraper interface {
	GetByKeywords(keywords string) (error, []article.Article)
}
