package webservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shops-scraping/database"
	"shops-scraping/shared"
	"strconv"
)

func addToCart(_ http.ResponseWriter, req *http.Request) httpResponse {
	var dto createCartLineDto
	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		return httpResponse{"dto not valid", http.StatusBadRequest}
	}

	line, err := database.CreateCartLine(shared.Article{
		Shop:       dto.Shop,
		Name:       dto.Name,
		Price:      dto.Price,
		Currency:   dto.Currency,
		Image:      dto.Image,
		DetailsUrl: dto.DetailsUrl,
	})
	if err != nil {
		return httpResponse{"Error during cart line save", http.StatusInternalServerError}
	}

	return httpResponse{line, http.StatusOK}
}

func getCart(_ http.ResponseWriter, _ *http.Request) httpResponse {
	cartLines, err := database.GetCartLines()
	if err != nil {
		return httpResponse{"Error during cart getting", http.StatusInternalServerError}
	}

	var lines []cartLineDto
	for _, line := range cartLines {
		lines = append(lines, newCartLineDto(line))
	}

	if len(lines) == 0 {
		return httpResponse{cartLines, http.StatusOK}
	}
	return httpResponse{lines, http.StatusOK}

}

func clearCart(_ http.ResponseWriter, _ *http.Request) httpResponse {
	err := database.ClearCart()
	if err != nil {
		return httpResponse{"Error during cart clear", http.StatusInternalServerError}
	}

	return httpResponse{"Cart clear successfully", http.StatusOK}
}

func deleteCartLine(_ http.ResponseWriter, r *http.Request) httpResponse {
	idParam := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return httpResponse{"id must be required", http.StatusBadRequest}
	}

	err = database.DeleteCartLine(id)
	if err != nil {
		return httpResponse{"Error during cart getting", http.StatusInternalServerError}
	}

	return httpResponse{fmt.Sprintf("line with id %s deleted successfully", idParam), http.StatusOK}
}
