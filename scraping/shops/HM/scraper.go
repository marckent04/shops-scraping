package HM

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"shops-scraping/scraping/Browser"
	"shops-scraping/scraping/common"
	"shops-scraping/shared"
	"sync"
)

type Scraper struct {
	url string
}

func (s Scraper) GetByKeywords(p common.SearchParams) ([]shared.Article, error) {
	msg := fmt.Sprintf("Search for %s at %s for %s", p.Keywords, shopName, p.Gender)
	log.Println(msg, " started")
	defer log.Println(msg, " finished")

	browser := Browser.GetInstance()

	page := browser.MustPage(fmt.Sprintf("%s?q=%s&department=%s", searchUrl, p.Keywords, genders[p.Gender]))
	defer page.MustClose()

	err := page.WaitElementsMoreThan(productSelector, 0)
	if err != nil {
		return make([]shared.Article, 0), err
	}

	if !page.MustHas(productsListSelector) {
		return make([]shared.Article, 0), nil
	}
	common.CloseCookieDialog(page)

	foundArticles := page.MustElements(productSelector)

	collection := common.ArticlesCollection{}
	var wg sync.WaitGroup
	for _, v := range foundArticles {
		node := v
		node = node.MustHover()
		wg.Add(1)
		go func() {
			defer wg.Done()
			collection.Push(rodeToArticle(node))
		}()
	}

	wg.Wait()

	return collection.Get(), nil
}

func NewScrapper() common.Scraper {
	return &Scraper{url: searchUrl}
}
