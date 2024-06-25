package SHEIN

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"golang.org/x/net/context"
	"golang.org/x/net/html"
	"log"
	"shops-scraping/scraping/common"
	common2 "shops-scraping/shared"
	"strconv"
	"strings"
)

func getProducts(keyword string) (items []common2.Article) {
	ch := make(chan []common2.Article, 1)

	go getProductsWithCaptcha(ch, keyword)
	go getProductsWithoutCaptcha(ch, keyword)

	items = <-ch

	return
}

func getProductsWithoutCaptcha(ch chan<- []common2.Article, keywords string) (err error) {
	var articles []common2.Article

	ctx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		chromedp.Headless,
		chromedp.NoSandbox,
	)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var h string

	log.Println("shein products getting in progress ...")

	err = chromedp.Run(ctx,
		network.SetExtraHTTPHeaders(common.Headers),
		chromedp.Navigate(fmt.Sprintf("%s/%s", searchUrl, keywords)),
		chromedp.InnerHTML(productsListSelector, &h, chromedp.ByQuery),
	)

	log.Println("shein products getting in finished")
	if err != nil {
		log.Println(err.Error())
	}

	document, err := goquery.NewDocumentFromReader(strings.NewReader(h))

	if err != nil {
		return err
	}

	for _, node := range document.Find(productSelector).Nodes {
		articles = append(articles, nodeToArticle(node))
	}

	ch <- articles
	return
}

func getProductsWithCaptcha(ch chan<- []common2.Article, keywords string) (err error) {
	var articles []common2.Article

	ctx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		chromedp.Headless,
		chromedp.NoSandbox,
	)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)

	defer cancel()

	var h string

	log.Println("shein products getting expected captcha in progress ...")

	err = chromedp.Run(ctx,
		network.SetExtraHTTPHeaders(common.Headers),
		chromedp.Navigate(fmt.Sprintf("%s/%s", searchUrl, keywords)),
		chromedp.WaitVisible(".geetest_close", chromedp.ByQuery),
		chromedp.Click(".geetest_close"),
		chromedp.WaitNotVisible(".geetest_close"),
		chromedp.SetJavascriptAttribute("input[name=\"header-search\"]", "value", keywords),
		chromedp.Click("form.header-search .search-btn", chromedp.ByQuery),
		chromedp.InnerHTML(productsListSelector, &h, chromedp.ByQuery),
	)

	log.Println("shein products getting expected captcha finished")
	if err != nil {
		return err
	}

	document, err := goquery.NewDocumentFromReader(strings.NewReader(h))
	if err != nil {
		return err
	}

	for _, node := range document.Find(productSelector).Nodes {
		articles = append(articles, nodeToArticle(node))
	}

	ch <- articles
	return
}

func nodeToArticle(node *html.Node) common2.Article {
	doc := goquery.NewDocumentFromNode(node)

	name, image, detailsPath, price :=
		getProductName(doc),
		strings.Replace(common.GetAttrValue(doc, "img", "src"), "//", "", 1),
		common.GetAttrValue(doc, "a.goods-title-link", "href"),
		getProductPrice(doc)

	return common2.New(name, image, fmt.Sprintf("%s%s", baseUrl, detailsPath), "SHEIN", price, "€")
}

func getProductPrice(doc *goquery.Document) float32 {
	priceTxt := doc.Find(".product-card__prices-info").Text()
	priceTxt = strings.Replace(priceTxt, ",", ".", 1)

	priceTxt, exists := strings.CutSuffix(priceTxt, "€")

	if !exists {
		return 0
	}

	price, err := strconv.ParseFloat(priceTxt, 32)

	if err != nil {
		log.Printf("Error during H&M product price getting: %v", err)
		return 0
	}

	return float32(price)
}

func getProductName(doc *goquery.Document) string {
	title := strings.Replace(doc.Find("a.goods-title-link").Text(), "SHEIN", "", 1)
	return strings.Trim(title, " ")
}
