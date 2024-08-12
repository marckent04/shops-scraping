package PULLBEAR

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
	log.Printf("%s products getting in progress ...", shopName)

	var articles []shared.Article

	browser := Browser.GetInstance()

	page := browser.MustPage(fmt.Sprintf(searchUrl, genders[p.Gender], p.Keywords)).MustWaitDOMStable()
	defer page.MustClose()
	defer log.Printf("%s products getting finished ...", shopName)

	page.MustWaitElementsMoreThan("search-app", 0)
	grid := getSearchGrid(page)
	common.CloseCookieDialog(page)

	if !grid.MustHas(".results") {
		return articles, nil
	}

	collection := common.ArticlesCollection{}
	var wg sync.WaitGroup

	foundArticles := getArticlesNodes(grid)
	for _, art := range foundArticles {
		node := art
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

	articles = append(articles, collection.Get()...)

	return articles, nil
}

func NewScrapper() common.Scraper {
	return &Scraper{url: searchUrl}
}
