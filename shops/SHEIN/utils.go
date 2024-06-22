package SHEIN

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"golang.org/x/net/context"
	"golang.org/x/net/html"
	"log"
	"shops-scraping/common"
	"strconv"
	"strings"
)

func getProducts(keywords string) (err error, articles []common.Article) {
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
		network.SetExtraHTTPHeaders(map[string]interface{}{
			"Accept":         "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
			"Sec-Fetch-Dest": "document",
			"Sec-Fetch-Mode": "navigate",
			"Sec-Fetch-Site": "none",
			"User-Agent":     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4 Safari/605.1.15",
		}),
		chromedp.Navigate(fmt.Sprintf("%s/%s", searchUrl, keywords)),
		chromedp.WaitVisible(".geetest_close", chromedp.ByQuery),
		chromedp.Click(".geetest_close"),
		chromedp.WaitNotVisible(".geetest_close"),
		chromedp.SetJavascriptAttribute("input[name=\"header-search\"]", "value", keywords),
		chromedp.Click("form.header-search .search-btn", chromedp.ByQuery),
		chromedp.InnerHTML(productsListSelector, &h, chromedp.ByQuery),
	)

	log.Println("shein products getting in finished")
	if err != nil {
		log.Println("laa")
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

func nodeToArticle(node *html.Node) common.Article {
	doc := goquery.NewDocumentFromNode(node)

	name, image, detailsPath, price :=
		doc.Find("a.goods-title-link").Text(),
		strings.Replace(common.GetAttrValue(doc, "img", "src"), "//", "", 1),
		common.GetAttrValue(doc, "a.goods-title-link", "href"),
		getProductPrice(doc)

	return common.New(name, image, fmt.Sprintf("%s%s", baseUrl, detailsPath), "SHEIN", price, "€")
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
