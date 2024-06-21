package HM

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"log"
	"shops-scraping/article"
	"shops-scraping/common"
	"shops-scraping/scraping"
	"strings"
)

type Scraper struct {
	url string
}

func (s Scraper) GetByKeywords(keywords string) (err error, articles []article.Article) {
	err, articles = s.getProducts(keywords)
	if err != nil {
		return err, nil
	}

	return
}

func (s Scraper) getProducts(keywords string) (err error, articles []article.Article) {
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
		chromedp.Navigate(fmt.Sprintf("%s?q=%s", s.url, keywords)),
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

func NewScrapper() scraping.Scraper {
	return &Scraper{url: searchUrl}
}
