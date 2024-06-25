package shared

type Article struct {
	Shop       string  `json:"shop"`
	Name       string  `json:"title"`
	Price      float32 `json:"price"`
	Currency   string  `json:"currency"`
	Image      string  `json:"image"`
	DetailsUrl string  `json:"detailsUrl"`
}

func New(name, image, detailsUrl, shop string, price float32, currency string) Article {
	return Article{
		Name:       name,
		Image:      image,
		Price:      price,
		DetailsUrl: detailsUrl,
		Currency:   currency,
		Shop:       shop,
	}
}
