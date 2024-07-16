package BERSHKA

import (
	"shops-scraping/scraping/common"
	"shops-scraping/shared"
	"strings"

	"github.com/go-rod/rod"
)

func rodeToArticle(elt *rod.Element) shared.Article {
	name, _ := elt.MustElement("[data-qa-anchor=\"productItemText\"]").Text()
	image, _ := elt.MustElement("img").Attribute("data-original")
	detailsUrl, _ := elt.MustElement(".grid-card-link").Attribute("href")
	priceText, _ := elt.MustElement("[data-qa-anchor=\"productItemPrice\"]").Text()
	price := getProductPrice(priceText)

	return shared.New(name, *image, *detailsUrl, "H&M", price, "€")
}

func getProductPrice(priceText string) float32 {
	priceLabel := strings.Replace(priceText, ",", ".", 1)
	return common.GetPrice(priceLabel)
}
