package webservice

import (
	"shops-scraping/database"
	"shops-scraping/shared"
	"time"
)

type createCartLineDto struct {
	shared.Article
}

type cartLineDto struct {
	shared.Article
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

func newCartLineDto(line database.CartLine) cartLineDto {
	return cartLineDto{
		Article: shared.Article{
			Shop:       line.Shop,
			DetailsUrl: line.DetailsUrl,
			Price:      line.Price,
			Name:       line.Name,
			Image:      line.Image,
			Currency:   line.Currency,
		},
		ID:        line.ID,
		CreatedAt: line.CreatedAt,
	}
}
