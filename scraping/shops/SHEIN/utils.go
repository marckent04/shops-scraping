package SHEIN

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"shops-scraping/shared"
	"strconv"
	"strings"
	"time"
)

func getProducts(keyword string) (items []shared.Article) {

	log.Info("Shein articles scrapping begins ...")
	ch := make(chan []shared.Article, 1)

	url, err := launcher.New().Headless(false).Devtools(false).Launch()

	page := rod.New().NoDefaultDevice().ControlURL(url).MustConnect().MustIncognito().MustPage(baseUrl).MustWaitLoad()
	defer page.MustClose()

	//page = page.MustWaitDOMStable()

	log.Println(page.MustHTML())
	var articles []shared.Article

	//ctx, _ := context.WithCancel(context.Background())

	//page = page.Context(ctx)

	captchaBtn := ".geetest_close"

	page = page.Timeout(2 * time.Second)

	waitErr := page.WaitElementsMoreThan(captchaBtn, 0)

	if waitErr == nil {
		log.Info("captcha resolved")
		clickErr := page.MustElement(captchaBtn).MustClick()
		if clickErr != nil {
			log.Error("close captcha error", err)
		}
	} else {
		log.Info("no captcha btn")
	}

	discountCloseBtn := ".coupon-dialog .dialog-header-v2__close-btn .btn-default"
	waitErr = page.WaitElementsMoreThan(discountCloseBtn, 0)
	if waitErr == nil {
		clickErr := page.MustElement(discountCloseBtn).MustClick()
		if clickErr != nil {
			log.Error("close discount error", err)
			return
		} else {
			log.Info("discount closed")
		}
	}

	log.Info("get search input")

	sErr := page.MustElement("form.header-search input").Input(keyword)
	if sErr != nil {
		log.Error("Error during search: ", sErr)
	}

	log.Info("submit form")

	clickErr := page.MustElement("form.header-search .search-btn").Click(proto.InputMouseButtonLeft, 1)

	if clickErr != nil {
		log.Println("submit form error", clickErr)
	} else {
		log.Info("search launched")
	}

	log.Info("wait for products")

	err = page.WaitElementsMoreThan(productSelector, 4)

	if err != nil {
		log.Error("elements not found", err)
		return
	}

	log.Info("elements found")

	articlesNodes := page.MustElements(productSelector)

	log.Info(articlesNodes)

	for i, node := range articlesNodes {
		articles = append(articles, rodeToArticle(*node))
		log.Infof("article %d/%d processed", i, len(articlesNodes))
	}

	log.Info("articles getting finished")

	items = <-ch

	return
}

func getOnlyArticles(page *rod.Page, keyword string) {
	inputErr := page.MustElement("input[name=\"header-search\"]").Input(keyword)

	if inputErr != nil {
		log.Println("inputErr", inputErr)
	}

	clickErr := page.MustElement("form.header-search .search-btn").Click(proto.InputMouseButtonLeft, 1)

	if clickErr != nil {
		log.Println("submit form error", clickErr)
	}

	page = page.MustWaitElementsMoreThan(productSelector, 4)

	articlesNodes := page.MustElements(productSelector)

	log.Println(articlesNodes)

}

func getProductsWithCaptcha(pageValue rod.Page, ch chan<- []shared.Article, keyword string) (err error) {
	var articles []shared.Article

	ctx, _ := context.WithCancel(context.Background())
	page := &pageValue

	page = page.Context(ctx)

	captchaBtn := ".geetest_close"

	waitErr := page.Timeout(2*time.Second).WaitElementsMoreThan(captchaBtn, 0)

	if waitErr == nil {
		log.Info("captcha resolved")
		clickErr := page.MustElement(captchaBtn).Click(proto.InputMouseButtonLeft, 1)
		if clickErr != nil {
			log.Error("close captcha error", err)
			return
		}
	} else {
		log.Error("error during captcha waiting ", waitErr)
	}

	log.Info("get search input")
	sErr := page.MustElement("input[name=\"header-search\"]").Input(keyword)
	if sErr != nil {
		log.Error("Error during search: ", sErr)
	}

	log.Info("submit form")

	clickErr := page.MustElement("form.header-search .search-btn").Click(proto.InputMouseButtonLeft, 1)

	if clickErr != nil {
		log.Println("submit form error", clickErr)
	} else {
		log.Info("search launched")
	}

	log.Info("wait for products")

	err = page.WaitElementsMoreThan(productSelector, 4)

	if err != nil {
		log.Error("elements not found", err)
		return
	}

	log.Info("elements found")

	articlesNodes := page.MustElements(productSelector)

	log.Info(articlesNodes)

	for i, node := range articlesNodes {
		articles = append(articles, rodeToArticle(*node))
		log.Infof("article %d/%d processed", i, len(articlesNodes))
	}

	log.Info("articles getting finished")

	//page = page.MustWaitElementsMoreThan(productSelector, 4)

	// ctx, cancel := chromedp.NewExecAllocator(
	// 	context.Background(),
	// 	chromedp.Headless,
	// 	chromedp.NoSandbox,
	// )
	// defer cancel()

	// ctx, cancel = chromedp.NewContext(ctx)

	// defer cancel()

	// var h string

	// log.Println("shein products getting expected captcha in progress ...")

	// err = chromedp.Run(ctx,
	// 	network.SetExtraHTTPHeaders(common.Headers),
	// 	chromedp.Navigate(fmt.Sprintf("%s/%s", searchUrl, keywords)),
	// 	chromedp.WaitVisible(".geetest_close", chromedp.ByQuery),
	// 	chromedp.Click(".geetest_close"),
	// 	chromedp.WaitNotVisible(".geetest_close"),
	// 	chromedp.SetJavascriptAttribute("input[name=\"header-search\"]", "value", keywords),
	// 	chromedp.Click("form.header-search .search-btn", chromedp.ByQuery),
	// 	chromedp.InnerHTML(productsListSelector, &h, chromedp.ByQuery),
	// )

	// log.Println("shein products getting expected captcha finished")
	// if err != nil {
	// 	return err
	// }

	// document, err := goquery.NewDocumentFromReader(strings.NewReader(h))
	// if err != nil {
	// 	return err
	// }

	// for _, node := range document.Find(productSelector).Nodes {
	// 	articles = append(articles, nodeToArticle(node))
	// }

	ch <- articles
	return
}

func rodeToArticle(node rod.Element) shared.Article {
	name, image, detailsPath, price :=
		getProductName(node),
		"",
		//getArticleImage(node),
		"",
		//*node.MustElement("a.goods-title-link").MustAttribute("href"),
		float32(0)
	//getProductPrice(node)

	return shared.New(name, image, fmt.Sprintf("%s%s", baseUrl, detailsPath), "SHEIN", price, "€")
}

func getProductPrice(doc rod.Element) float32 {
	priceElt, err := doc.Element(".product-card__prices-info span")

	if err != nil {
		return 0
	}

	priceTxt := priceElt.MustText()

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

func getProductName(doc rod.Element) string {
	title, err := doc.MustElement("a.goods-title-link").Text()
	if err != nil {
		return ""
	}
	return strings.Trim(strings.Replace(title, "SHEIN", "", 1), " ")
}

func getArticleImage(doc rod.Element) string {
	value, err := doc.MustElement(".crop-image-container__inner").Attribute("style")

	if err != nil {
		log.Println("no image")
		return "ok"
	}
	if len(*value) > 1 {
		return extractImage(*value)
	}

	return *doc.MustElement(".crop-image-container__img").MustAttribute("src")
}

func extractImage(style string) string {
	b, e := strings.Index(style, "("),
		strings.Index(style, ")")

	return style[b+1 : e]
}
