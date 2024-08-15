package webservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shops-scraping/database"
	"shops-scraping/shared"
	"strconv"
)

func addToCart(rsp http.ResponseWriter, req *http.Request) {
	var dto createCartLineDto
	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		serveMessageResponse(rsp, "dto not valid", http.StatusBadRequest)
		return
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
		serveMessageResponse(rsp, "Error during cart line save", http.StatusInternalServerError)
		return
	}

	serveJsonResponse(rsp, line)
}

func getCart(rsp http.ResponseWriter, _ *http.Request) {
	cartLines, err := database.GetCartLines()
	if err != nil {
		serveMessageResponse(rsp, "Error during cart getting", http.StatusInternalServerError)
		return
	}

	var lines []cartLineDto
	for _, line := range cartLines {
		lines = append(lines, newCartLineDto(line))
	}

	if len(lines) == 0 {
		serveJsonResponse(rsp, cartLines)
		return
	}
	serveJsonResponse(rsp, lines)
}

func clearCart(rsp http.ResponseWriter, _ *http.Request) {
	err := database.ClearCart()
	if err != nil {
		serveMessageResponse(rsp, "Error during cart clear", http.StatusInternalServerError)
		return
	}

	serveMessageResponse(rsp, "Cart clear successfully", http.StatusOK)
}

func deleteCartLine(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		serveMessageResponse(w, "id must be required", http.StatusBadRequest)
		return
	}

	err = database.DeleteCartLine(id)
	if err != nil {
		serveMessageResponse(w, "Error during cart getting", http.StatusInternalServerError)
		return
	}

	serveMessageResponse(w, fmt.Sprintf("line with id %s deleted successfully", idParam), http.StatusOK)
}
