package webservice

import (
	"net/http"
	"shops-scraping/scraping/shops"
)

func getEnabledShops(_ http.ResponseWriter, _ *http.Request) httpResponse {
	var dtos []shopDto

	for _, shop := range shops.GetEnabledShops() {
		dtos = append(dtos, newShopDto(shop.Name, shop.Code))
	}

	return httpResponse{dtos, http.StatusOK}
}
