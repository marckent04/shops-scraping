package BERSHKA

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"shops-scraping/scraping/common"
	"shops-scraping/shared"
	"strings"

	"github.com/go-rod/rod"
)

func getProducts(keyword string) (err error, articles []shared.Article) {
	log.Printf("%s products getting in progress ...", shopName)

	page := rod.New().MustConnect().MustPage(fmt.Sprintf("%s/%s", searchUrl, keyword)).MustWaitDOMStable()

	isEmpty := page.MustHas(".search-empty__empty")

	if isEmpty {
		return
	}

	page = page.MustWaitElementsMoreThan(articleSelector, 4)

	page.Mouse.Scroll(10, 10000, 15)

	foundArticles := page.MustElements(articleSelector)

	for _, node := range foundArticles {
		h, _ := node.HTML()

		isArticle := strings.Contains(h, "data-qa-anchor=\"productItemText\"")

		if !isArticle {
			continue
		}

		articles = append(articles, rodeToArticle(node))
	}

	log.Printf("%s articles getting finished", shopName)

	if err != nil {
		log.Println(err.Error())
	}

	return
}

func rodeToArticle(elt *rod.Element) shared.Article {
	name, _ := elt.MustElement("[data-qa-anchor=\"productItemText\"]").Text()
	image, _ := elt.MustElement("img").Attribute("data-original")
	detailsUrl, _ := elt.MustElement(".grid-card-link").Attribute("href")
	pricetext, _ := elt.MustElement("[data-qa-anchor=\"productItemPrice\"]").Text()
	price := getProductPrice(pricetext)

	return shared.New(name, *image, *detailsUrl, "H&M", price, "€")
}

func getProductPrice(priceText string) float32 {
	priceLabel := strings.Replace(priceText, ",", ".", 1)
	return common.GetPrice(priceLabel)
}
