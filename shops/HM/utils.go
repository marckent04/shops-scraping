package HM

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"shops-scraping/article"
	"shops-scraping/common"
	"strconv"
	"strings"
)

func nodeToArticle(node *html.Node) article.Article {
	document := goquery.NewDocumentFromNode(node)

	name, image, detailsUrl :=
		common.GetAttrValue(document, "a", "title"),
		common.GetAttrValue(document, "img", "src"),
		common.GetAttrValue(document, "a", "href")

	price := document.Find("span").FilterFunction(func(i int, selection *goquery.Selection) bool {
		return strings.Contains(selection.Text(), "€")
	}).First().Text()

	return newArticle(name, image, price, detailsUrl)
}

func newArticle(name, image, priceWithCurrency, detailsUrl string) article.Article {
	priceLabel := strings.Replace(priceWithCurrency, ",", ".", 1)
	values := strings.Split(priceLabel, " ")
	price, err := strconv.ParseFloat(values[0], 32)

	if err != nil {
		return article.Article{}
	}

	return article.Article{
		Name:       name,
		Image:      image,
		Price:      float32(price),
		DetailsUrl: detailsUrl,
		Currency:   values[1],
		Shop:       "H&M",
	}
}
