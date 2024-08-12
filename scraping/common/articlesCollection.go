package common

import (
	"shops-scraping/shared"
	"sync"
)

type ArticlesCollection struct {
	mu    sync.Mutex
	items []shared.Article
}

func (c *ArticlesCollection) Push(article shared.Article) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = append(c.items, article)
}

func (c *ArticlesCollection) Get() []shared.Article {
	return c.items
}
