package SHEIN

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
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
		chromedp.NoDefaultBrowserCheck,
		//chromedp.UserAgent("Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Mobile Safari/537.36"),
		//chromedp.Headless,
		//chromedp.NoSandbox,
	)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var h string

	log.Println("shein products getting in progress ...")
	err = chromedp.Run(ctx,
		//network.SetExtraHTTPHeaders(common.GetBrowserHeaders()),
		chromedp.Navigate(fmt.Sprintf("%s/%s", searchUrl, keywords)),
		chromedp.WaitVisible(".geetest_close", chromedp.ByQuery),
		chromedp.Click(".geetest_close"),
		chromedp.WaitNotVisible(".geetest_close"),
		chromedp.SetJavascriptAttribute("input[name=\"header-search\"]", "value", keywords),
		chromedp.Click("form.header-search .search-btn", chromedp.ByQuery),
		chromedp.WaitReady(productsListSelector, chromedp.ByQuery),
		chromedp.InnerHTML(productsListSelector, &h, chromedp.ByQuery),
	)

	log.Println("shein products getting in finished")

	log.Println(h)
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

	name, image, detailsPath, price :=
		doc.Find("a.goods-title-link").Text(),
		strings.Replace(common.GetAttrValue(doc, "img", "src"), "//", "", 1),
		common.GetAttrValue(doc, "a.goods-title-link", "href"),
		getProductPrice(doc)

	return article.New(name, image, fmt.Sprintf("%s%s", baseUrl, detailsPath), "H&M", price, "€")
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
