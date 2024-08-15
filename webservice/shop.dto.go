package webservice

type shopDto struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

func newShopDto(name, code string) shopDto {
	return shopDto{
		Text:  name,
		Value: code,
	}
}
