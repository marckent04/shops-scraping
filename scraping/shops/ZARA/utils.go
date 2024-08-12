package ZARA

import (
	log "github.com/sirupsen/logrus"
	"shops-scraping/shared"
	"strconv"
	"strings"

	"github.com/go-rod/rod"
)

func rodeToArticle(elt *rod.Element) (article shared.Article, err error) {
	price, err := getProductPrice(elt)
	if err != nil {
		return
	}

	name := elt.MustElement(".product-grid-product-info__main-info").MustText()
	image := getArticleImg(elt)
	detailsUrl, _ := elt.MustElement(".product-link").Attribute("href")

	article = shared.New(strings.ToTitle(name), image, *detailsUrl, "ZARA", price, "€")
	return
}

func getArticleImg(elt *rod.Element) string {
	return *elt.MustElement(".media-image__image").MustAttribute("src")
}

func getProductPrice(elt *rod.Element) (float32, error) {
	prices, err := elt.Elements(".money-amount__main")
	if err != nil {
		return 0, err
	}

	priceNode := prices.Last()
	priceText := priceNode.MustText()

	priceLabel := strings.Replace(priceText, ",", ".", 1)
	results := strings.Split(priceLabel, " ")
	price, err := strconv.ParseFloat(results[0], 32)
	if err != nil {
		log.Printf("Error during ZARA product price getting: %v", err)
		return 0, err
	}

	return float32(price), nil
}
