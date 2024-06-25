package BERSHKA

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"golang.org/x/net/context"
	"golang.org/x/net/html"
	"log"
	"regexp"
	"shops-scraping/scraping/common"
	common2 "shops-scraping/shared"
	"strconv"
	"strings"
)

func getProducts(keyword string) (err error, articles []common2.Article) {
	ctx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		chromedp.Headless,
		chromedp.NoSandbox,
	)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var h string

	log.Printf("%s products getting in progress ...", shopName)

	//var pic []byte
	err = chromedp.Run(ctx,
		network.SetExtraHTTPHeaders(network.Headers{
			"Accept":           "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
			"Accept-Encoding":  "gzip, deflate, br, zstd",
			"Accept-Language":  "fr-CI,fr;q=0.9,en-GB;q=0.8,en;q=0.7,fr-FR;q=0.6,en-US;q=0.5",
			"Cache-Control":    "max-age=0",
			"f-None-Match":     "107tr9n2kobde93",
			"Priority":         "u=0, i",
			"Sec-Ch-Ua":        "\"Not/A)Brand\";v=\"8\", \"Chromium\";v=\"126\", \"Google Chrome\";v=\"126\"",
			"Sec-Ch-Ua-Mobile": "?1", "Sec-Ch-Ua-Platform": "Android",
			"Sec-Fetch-Dest":            "document",
			"Sec-Fetch-Mode":            "navigate",
			"Sec-Fetch-Site":            "same-origin",
			"Sec-Fetch-User":            "?1",
			"Upgrade-Insecure-Requests": "1",
			"User-Agent":                "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Mobile Safari/537.36",
		}),
		chromedp.Navigate(fmt.Sprintf("%s/%s", searchUrl, keyword)),
		chromedp.WaitReady(articleSelector, chromedp.ByQuery),
		chromedp.InnerHTML(articlesListSelector, &h),
	)

	//ioutil.WriteFile("file.png", pic, 0666)

	log.Printf("%s articles getting finished", shopName)

	if err != nil {
		log.Println(err.Error())
	}

	document, err := goquery.NewDocumentFromReader(strings.NewReader(h))
	if err != nil {
		return err, nil
	}

	for _, node := range document.Find(articleSelector).Nodes {
		articles = append(articles, nodeToArticle(node))
	}

	fmt.Println(articles)
	return
}

func nodeToArticle(node *html.Node) common2.Article {
	doc := goquery.NewDocumentFromNode(node)

	name, image, detailsUrl, price :=
		doc.Find("[data-qa-anchor=\"productItemText\"]").Text(),
		common.GetAttrValue(doc, ".image-item", "src"),
		common.GetAttrValue(doc, ".grid-card-link", "href"),
		getProductPrice(doc.Find("[data-qa-anchor=\"productItemPrice\"]").Text())

	return common2.New(name, image, detailsUrl, "H&M", price, "€")
}

func getProductPrice(priceText string) float32 {
	priceLabel := strings.Replace(priceText, ",", ".", 1)

	re := regexp.MustCompile("\\d+\\.\\d+")
	results := re.FindStringSubmatch(priceLabel)

	if len(results) == 0 {
		return 0
	}
	price, err := strconv.ParseFloat(results[0], 32)

	if err != nil {
		log.Printf("Error during H&M product price getting: %v", err)
		return 0
	}

	return float32(price)
}
