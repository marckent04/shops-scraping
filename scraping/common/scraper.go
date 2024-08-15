package common

import (
	"os"
	"shops-scraping/shared"
)

type Gender = string

const (
	MAN   Gender = "m"
	WOMAN        = "w"
)

type Scraper interface {
	GetByKeywords(params SearchParams) ([]shared.Article, error)
}

type SearchParams struct {
	Gender   Gender
	Keywords string
}

func NewSearchParams(gender Gender, keywords string) SearchParams {
	return SearchParams{
		Gender:   gender,
		Keywords: keywords,
	}
}

type DisabledScraper struct {
}

func (_ DisabledScraper) GetByKeywords(_ SearchParams) (art []shared.Article, err error) {
	return
}

func NewScraper(fFVar string, instanceHandler func() Scraper) Scraper {
	if os.Getenv(fFVar) == "true" {
		return instanceHandler()
	}
	return DisabledScraper{}
}
