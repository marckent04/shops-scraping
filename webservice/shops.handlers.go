package webservice

import (
	"net/http"
	"shops-scraping/scraping/shops"
)

func getEnabledShops(w http.ResponseWriter, _ *http.Request) {
	var dtos []shopDto

	for _, shop := range shops.GetEnabledShops() {
		dtos = append(dtos, newShopDto(shop.Name, shop.Code))
	}
	serveJsonResponse(w, dtos)
}
