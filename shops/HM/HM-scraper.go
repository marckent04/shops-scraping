package HM

import (
	"fmt"
	"github.com/gocolly/colly"
	"shop-scraping/article"
	"shop-scraping/scraping"
	"strconv"
	"strings"
)

type Scraper struct {
	url string
}

func (s Scraper) GetByKeywords(keywords string) (err error, articles []article.Article) {

	c := colly.NewCollector(
		colly.MaxDepth(5),
		colly.IgnoreRobotsTxt(),
		colly.AllowURLRevisit(),
		colly.MaxBodySize(99999999),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36"),
	)

	c.OnHTML("#main-content > section > ul > li", func(product *colly.HTMLElement) {
		var price string
		image := product.ChildAttr("img", "src")
		name := product.ChildAttr("a", "title")
		details := product.ChildAttr("a", "href")
		product.ForEachWithBreak("span", func(_ int, element *colly.HTMLElement) bool {
			if strings.HasSuffix(element.Text, "€") {
				price = element.Text
				return false
			}
			return true
		})

		articles = append(articles, NewArticle(name, image, price, details))

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(response *colly.Response) {
		body := string(response.Body)

	})

	err = c.Visit(fmt.Sprintf("%s?q=%s&isE4=false", s.url, keywords))

	//log.Printf("total produits:  %d", len(articles))

	return
}

func NewScrapper() scraping.Scraper {
	return &Scraper{url: "https://www2.hm.com/fr_fr/search-results.html"}
}

func NewArticle(name, image, priceAndCurrency, detailsUrl string) article.Article {
	priceLabel := strings.Replace(priceAndCurrency, ",", ".", 1)
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
