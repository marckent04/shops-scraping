package article

type Article struct {
	Shop       string  `json:"shop"`
	Name       string  `json:"title"`
	Price      float32 `json:"price"`
	Currency   string  `json:"currency"`
	Image      string  `json:"image"`
	DetailsUrl string  `json:"detailsUrl"`
}
