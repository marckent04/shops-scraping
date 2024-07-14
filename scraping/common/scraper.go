package common

import (
	"github.com/go-rod/rod"
	"shops-scraping/shared"
)

type Scraper interface {
	GetByKeywords(browser *rod.Browser, keywords string) (error, []shared.Article)
}
