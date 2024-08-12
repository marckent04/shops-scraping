package cart

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shops-scraping/database"
	"shops-scraping/shared"
	"shops-scraping/webservice/utils"
	"strconv"
)

func AddToCart(rsp http.ResponseWriter, req *http.Request) {
	var dto createCartLineDto
	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		utils.ServeMessageResponse(rsp, "dto not valid", http.StatusBadRequest)
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
		utils.ServeMessageResponse(rsp, "Error during cart line save", http.StatusInternalServerError)
		return
	}

	utils.ServeJsonResponse(rsp, line)
}

func GetCart(rsp http.ResponseWriter, _ *http.Request) {
	cartLines, err := database.GetCartLines()
	if err != nil {
		utils.ServeMessageResponse(rsp, "Error during cart getting", http.StatusInternalServerError)
		return
	}

	var lines []cartLineDto
	for _, line := range cartLines {
		lines = append(lines, newCartLineDto(line))
	}

	if len(lines) == 0 {
		utils.ServeJsonResponse(rsp, cartLines)
		return
	}
	utils.ServeJsonResponse(rsp, lines)
}

func ClearCart(rsp http.ResponseWriter, _ *http.Request) {
	err := database.ClearCart()
	if err != nil {
		utils.ServeMessageResponse(rsp, "Error during cart clear", http.StatusInternalServerError)
		return
	}

	utils.ServeMessageResponse(rsp, "Cart clear successfully", http.StatusOK)
}

func DeleteCartLine(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.ServeMessageResponse(w, "id must be required", http.StatusBadRequest)
		return
	}

	err = database.DeleteCartLine(id)
	if err != nil {
		utils.ServeMessageResponse(w, "Error during cart getting", http.StatusInternalServerError)
		return
	}

	utils.ServeMessageResponse(w, fmt.Sprintf("line with id %s deleted successfully", idParam), http.StatusOK)
}
