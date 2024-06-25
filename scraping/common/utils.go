package common

import "github.com/PuerkitoBio/goquery"

func GetAttrValue(document *goquery.Document, selector, attr string) string {
	val, exists := document.Find(selector).Attr(attr)
	if exists {
		return val
	}
	return ""
}

var Headers = map[string]interface{}{
	"Accept":         "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	"Sec-Fetch-Dest": "document",
	"Sec-Fetch-Mode": "navigate",
	"Sec-Fetch-Site": "none",
	"User-Agent":     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4 Safari/605.1.15",
}
