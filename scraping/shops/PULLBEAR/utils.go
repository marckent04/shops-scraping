package PULLBEAR

import (
	"fmt"
	"github.com/go-rod/rod"
	log "github.com/sirupsen/logrus"
	"shops-scraping/scraping/common"
	"shops-scraping/shared"
	"strings"
)

func getProducts(browser *rod.Browser, keyword string) (err error, articles []shared.Article) {

	channel := make(chan []shared.Article)

	go getArticlesFor(browser, man, channel, keyword)
	go getArticlesFor(browser, woman, channel, keyword)

	articles = append(articles, <-channel...)
	articles = append(articles, <-channel...)

	return
}

func getArticlesFor(browser *rod.Browser, cat category, articlesChan chan<- []shared.Article, keyword string) {
	var articles []shared.Article

	log.Printf("%s products getting in progress ...", shopName)

	page := browser.MustPage(fmt.Sprintf(searchUrl, cat, keyword)).MustWaitDOMStable()

	acceptCookiesBtn := "#onetrust-accept-btn-handler"
	page.WaitElementsMoreThan(acceptCookiesBtn, 0)
	page.MustElement(acceptCookiesBtn).MustClick()

	err := page.Mouse.Scroll(0, 6000, 10)
	if err != nil {
		log.Error("error when scroll", err)
	}

	grid := getSearchGrid(page)

	if grid.MustHas(".results") {
		foundArticles := getArticlesSD(grid)
		for _, node := range foundArticles {
			articles = append(articles, rodeToArticle(node))
		}
	}

	articlesChan <- articles

	log.Printf("%s products getting finished ...", shopName)

	return
}

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
