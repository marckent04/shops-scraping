package PULLBEAR

import (
	"github.com/go-rod/rod"
	"shops-scraping/scraping/common"
	"shops-scraping/shared"
	"strings"
)

func rodeToArticle(elt *rod.Element) (shared.Article, error) {
	name := elt.MustElement(".product-name").MustText()
	detailsUrl := elt.MustElement("a").MustAttribute("href")
	price := getArticlePrice(elt)
	image := getArticleImg(elt)

	return shared.New(strings.ToTitle(name), image, *detailsUrl, shopName, price, "€"), nil
}

func getArticleImg(elt *rod.Element) string {
	media := elt.MustElement("x-media").MustShadowRoot().MustElement("lazy-image")
	return *media.MustAttribute("src")
}

func getArticlePrice(elt *rod.Element) float32 {
	price := elt.MustElement(".current-price")

	priceTxt := price.MustText()

	if elt.MustHas(".discount-price") {
		priceTxt = elt.MustElement(".discount-price").MustText()
	}

	priceLabel := strings.Replace(priceTxt, ",", ".", 1)

	return common.GetPrice(priceLabel)
}

func getArticlesNodes(searchGrid *rod.Element) (elements rod.Elements) {
	for _, element := range searchGrid.MustElements(articleSelector) {
		elements = append(elements, element.MustShadowRoot())
	}
	return
}

func getSearchGrid(page *rod.Page) *rod.Element {
	return page.MustElement("search-app").MustShadowRoot().MustElement("search-grid").MustShadowRoot()
}
