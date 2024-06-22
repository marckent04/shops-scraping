package common

type Scraper interface {
	GetByKeywords(keywords string) (error, []Article)
}
