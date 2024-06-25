package BERSHKA

const baseUrl = "https://www.bershka.com"
const searchUrl = "https://www.bershka.com/fr/q"
const articlesListSelector = ".search-grid"
const articleSelector = ".search-product-card"
const shopName = "Bershka"

var headers = map[string]interface{}{
	"Accept":         "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	"Sec-Fetch-Dest": "document",
	"Sec-Fetch-Mode": "navigate",
	"Sec-Fetch-Site": "none",
	"User-Agent":     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4 Safari/605.1.15",
}
