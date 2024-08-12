package common

import (
	"shops-scraping/shared"
)

type Gender = string

const (
	MAN   Gender = "m"
	WOMAN        = "w"
)

type Scraper interface {
	GetByKeywords(params SearchParams) (error, []shared.Article)
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
