package ZARA

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"shops-scraping/shared"
	"strconv"
	"strings"

	"github.com/go-rod/rod"
)

func getProducts(keyword string) (err error, articles []shared.Article) {

	channel := make(chan []shared.Article, 2)

	go getArticlesFor(man, channel, keyword)
	go getArticlesFor(woman, channel, keyword)

	articles = append(articles, <-channel...)
	articles = append(articles, <-channel...)

	return
}

func getArticlesFor(cat category, articlesChan chan<- []shared.Article, keyword string) {
	var articles []shared.Article

	log.Printf("%s products for %s getting in progress ...", shopName, cat)

	page := rod.New().MustConnect().MustPage(fmt.Sprintf(searchUrl, keyword, cat)).MustWaitDOMStable()

	if page.MustHas(".zds-empty-state__title") {
		articlesChan <- articles
		return
	}
	page = page.MustWaitElementsMoreThan(articleSelector, 1)

	page.Mouse.Scroll(10, 10000, 30)

	foundArticles := page.MustElements(articleSelector)

	for _, node := range foundArticles {
		if !node.MustHas(".money-amount__main") {
			continue
		}

		articles = append(articles, rodeToArticle(node))
	}

	articlesChan <- articles

	return
}

func rodeToArticle(elt *rod.Element) shared.Article {
	name := elt.MustElement(".product-grid-product-info__main-info").MustText()
	image := getArticleImg(elt)
	detailsUrl, _ := elt.MustElement(".product-link").Attribute("href")
	price := getProductPrice(elt)

	return shared.New(strings.ToTitle(name), image, *detailsUrl, "ZARA", price, "€")
}

func getArticleImg(elt *rod.Element) string {
	return *elt.MustElement(".media-image__image").MustAttribute("src")
}

func getProductPrice(elt *rod.Element) float32 {
	prices := elt.MustElements(".money-amount__main")
	priceNode := prices[len(prices)-1]
	priceText := priceNode.MustText()

	priceLabel := strings.Replace(priceText, ",", ".", 1)
	results := strings.Split(priceLabel, " ")
	price, err := strconv.ParseFloat(results[0], 32)

	if err != nil {
		log.Printf("Error during ZARA product price getting: %v", err)
		return 0
	}

	return float32(price)
}
