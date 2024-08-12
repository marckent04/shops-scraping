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

func (s Scraper) GetByKeywords(p common.SearchParams) (error, []shared.Article) {
	msg := fmt.Sprintf("Search for %s at %s for %s", p.Keywords, shopName, p.Gender)
	log.Println(msg, " started")
	defer log.Println(msg, " finished")

	browser := Browser.GetInstance()

	page := browser.MustPage(fmt.Sprintf("%s?q=%s&department=%s", searchUrl, p.Keywords, genders[p.Gender]))
	defer page.MustClose()

	err := page.WaitElementsMoreThan(productSelector, 0)
	if err != nil {
		return err, make([]shared.Article, 0)
	}

	if !page.MustHas(productsListSelector) {
		return nil, make([]shared.Article, 0)
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

	return nil, collection.Get()
}

func NewScrapper() common.Scraper {
	return &Scraper{url: searchUrl}
}
