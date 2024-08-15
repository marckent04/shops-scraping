package ZARA

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
	gender := genders[p.Gender]
	log.Printf("%s products for %s getting in progress ...\n", shopName, gender)
	defer log.Printf("%s products for %s getting finished\n", shopName, gender)

	browser := Browser.GetInstance()
	page := browser.MustPage(fmt.Sprintf(searchUrl, p.Keywords, gender))
	defer page.MustClose()

	hasResults := common.WaitForLoad(page, articleSelector, ".zds-empty-state__title")
	if !hasResults {
		return make([]shared.Article, 0), nil
	}

	common.CloseCookieDialog(page)

	foundArticles := page.MustElements(articleSelector)
	collection := common.ArticlesCollection{}
	var wg sync.WaitGroup
	for _, element := range foundArticles {
		node := element
		node.MustHover()

		wg.Add(1)
		go func() {
			article, err := rodeToArticle(node)
			if err == nil {
				collection.Push(article)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return collection.Get(), nil
}

func GetScrapper() common.Scraper {
	return common.NewScraper("ENABLE_ZARA", func() common.Scraper {
		return &Scraper{url: searchUrl}
	})
}
