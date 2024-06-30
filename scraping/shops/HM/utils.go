package HM

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"shops-scraping/shared"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
)

func getProducts(keywords string) (err error, articles []shared.Article) {

	log.Println("H&M products getting in progress ...")

	page := rod.New().MustConnect().MustPage(fmt.Sprintf("%s?q=%s", searchUrl, keywords)).MustWaitDOMStable()

	if !page.MustHas(productsListSelector) {
		return
	}

	page = page.MustWaitElementsMoreThan(productSelector, 5)

	page.Mouse.Scroll(10, 10000, 30)

	foundArticles := page.MustElements(productSelector)

	log.Println("H&M products getting finished")

	for _, v := range foundArticles {
		articles = append(articles, rodeToArticle(v))
	}

	return
}

func rodeToArticle(node *rod.Element) shared.Article {

	link := node.MustElement("a")

	name, _ := link.Attribute("title")
	detailsUrl, _ := link.Attribute("href")

	image, price :=
		getArticleImage(node),
		getProductPrice(node)

	return shared.New(*name, image, *detailsUrl, "H&M", price, "€")
}

func getProductPrice(node *rod.Element) float32 {
	h, _ := node.HTML()
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(h))

	priceNode := doc.Find("span").FilterFunction(func(i int, selection *goquery.Selection) bool {
		return strings.Contains(selection.Text(), "€")
	}).First()

	priceLabel := strings.Replace(priceNode.Text(), ",", ".", 1)
	values := strings.Split(priceLabel, " ")
	price, err := strconv.ParseFloat(values[0], 32)

	if err != nil {
		log.Printf("Error during H&M product price getting: %v", err)
		return 0
	}

	return float32(price)
}

func getArticleImage(doc *rod.Element) string {
	img := doc.MustElement("img")

	mainImg, err :=
		img.Attribute("src")

	if err != nil {
		return ""
	}

	return *mainImg
}
