package common

import "github.com/PuerkitoBio/goquery"

func GetAttrValue(document *goquery.Document, selector, attr string) string {
	val, exists := document.Find(selector).Attr(attr)
	if exists {
		return val
	}
	return ""
}
