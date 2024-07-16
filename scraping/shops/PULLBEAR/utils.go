package PULLBEAR

import (
	"github.com/go-rod/rod"
	"shops-scraping/scraping/common"
	"shops-scraping/shared"
	"strings"
)

func rodeToArticle(elt *rod.Element) shared.Article {
	eltSh := elt.MustShadowRoot()
	eltSh.MustElement(".product-name").MustText()
	detailsUrl := eltSh.MustElement("a").MustAttribute("href")
	name := strings.ToTitle(eltSh.MustElement(".product-name").MustText())
	price := getArticlePrice(eltSh)
	image := getArticleImg(eltSh)

	return shared.New(strings.ToTitle(name), image, *detailsUrl, shopName, price, "€")
}

func getArticleImg(elt *rod.Element) string {
	imgContainer := elt.MustElement("x-media").MustShadowRoot().MustElement("lazy-image").MustShadowRoot()
	return *imgContainer.MustElement("img").MustAttribute("src")
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

func getArticlesSD(searchGrid *rod.Element) rod.Elements {
	return searchGrid.MustElements(articleSelector)
}

func getSearchGrid(page *rod.Page) *rod.Element {
	return page.MustElement("search-app").MustShadowRoot().MustElement("search-grid").MustShadowRoot()
}
