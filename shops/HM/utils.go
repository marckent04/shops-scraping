package HM

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"golang.org/x/net/context"
	"golang.org/x/net/html"
	"log"
	"shops-scraping/article"
	"shops-scraping/common"
	"strconv"
	"strings"
)

func getProducts(keywords string) (err error, articles []article.Article) {
	ctx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		chromedp.Headless,
		chromedp.NoSandbox,
	)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var h string

	err = chromedp.Run(ctx,
		network.SetExtraHTTPHeaders(common.GetBrowserHeaders()),
		chromedp.Navigate(fmt.Sprintf("%s?q=%s", searchUrl, keywords)),
		chromedp.WaitReady("li:nth-of-type(36)", chromedp.ByQuery),
		chromedp.InnerHTML(productsListSelector, &h),
	)

	if err != nil {
		log.Println(err.Error())
	}

	document, err := goquery.NewDocumentFromReader(strings.NewReader(h))
	if err != nil {
		return err, nil
	}

	for _, node := range document.Find(productSelector).Nodes {
		articles = append(articles, nodeToArticle(node))
	}

	return
}

func nodeToArticle(node *html.Node) article.Article {
	doc := goquery.NewDocumentFromNode(node)

	name, image, detailsUrl, price :=
		common.GetAttrValue(doc, "a", "title"),
		common.GetAttrValue(doc, "img", "src"),
		common.GetAttrValue(doc, "a", "href"),
		getProductPrice(doc)

	return article.New(name, image, detailsUrl, "H&M", price, "€")
}

func getProductPrice(doc *goquery.Document) float32 {
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
