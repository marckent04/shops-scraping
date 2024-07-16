package common

import (
	"github.com/go-rod/rod"
	"shops-scraping/shared"
)

type Gender = string

const (
	MAN   Gender = "m"
	WOMAN        = "w"
)

type Scraper interface {
	GetByKeywords(browser *rod.Browser, params SearchParams) (error, []shared.Article)
}

type SearchParams struct {
	Gender   Gender
	Keywords Gender
}

func NewSearchParams(gender Gender, keywords string) SearchParams {
	return SearchParams{
		Gender:   gender,
		Keywords: keywords,
	}
}
