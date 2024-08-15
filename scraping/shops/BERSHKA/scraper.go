package BERSHKA

import (
	"fmt"
	"github.com/go-rod/rod"
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
	defer log.Printf("%s articles getting finished", shopName)

	browser := Browser.GetInstance()
	page := browser.MustPage(fmt.Sprintf("%s/%s?gender=%s", searchUrl, p.Keywords, genders[p.Gender]))
	defer page.MustClose()

	hasResults := s.waitForLoad(page)
	if !hasResults {
		return make([]shared.Article, 0), nil
	}

	foundArticles := page.MustElements(articleSelector)
	collection := common.ArticlesCollection{}
	var wg sync.WaitGroup
	for _, art := range foundArticles {
		node := art
		go func() {
			wg.Add(1)
			defer wg.Done()
			article, err := rodeToArticle(node)
			if err == nil {
				collection.Push(article)
			}
		}()
	}

	wg.Wait()

	return collection.Get(), nil
}

func (s Scraper) waitForLoad(page *rod.Page) (hasResults bool) {
	var ch = make(chan bool, 1)

	go func() {
		err := page.WaitElementsMoreThan(articleSelector, 0)
		if err == nil {
			ch <- true
		}
	}()
	go func() {
		err := page.WaitElementsMoreThan(".search-empty__empty", 0)
		if err == nil {
			ch <- false
		}
	}()

	hasResults = <-ch

	return
}

func GetScrapper() common.Scraper {
	return common.NewScraper("ENABLE_BERSHKA", func() common.Scraper {
		return &Scraper{url: searchUrl}
	})
}
