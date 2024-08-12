package BERSHKA

import (
	"shops-scraping/scraping/common"
	"shops-scraping/shared"
	"strings"

	"github.com/go-rod/rod"
)

func rodeToArticle(elt *rod.Element) (article shared.Article, err error) {
	nameNode, err := elt.Element("[data-qa-anchor=\"productItemText\"]")
	if err != nil {
		return
	}

	name := nameNode.MustText()
	image, _ := elt.MustElement("img").Attribute("data-original")
	detailsUrl, _ := elt.MustElement(".grid-card-link").Attribute("href")
	priceText, _ := elt.MustElement("[data-qa-anchor=\"productItemPrice\"]").Text()
	price := getProductPrice(priceText)

	url := "https://www.bershka.com" + *detailsUrl

	return shared.New(name, *image, url, shopName, price, "€"), nil
}

func getProductPrice(priceText string) float32 {
	priceLabel := strings.Replace(priceText, ",", ".", 1)
	return common.GetPrice(priceLabel)
}
