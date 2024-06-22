package HM

import "github.com/chromedp/cdproto/network"

const productsListSelector = "#main-content > section > ul"
const searchUrl = "https://www2.hm.com/fr_fr/search-results.html"
const productSelector = "article[data-articlecode]"

func getBrowserHeaders() network.Headers {
	return network.Headers{
		"Accept":           "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"Accept-Encoding":  "gzip, deflate, br, zstd",
		"Accept-Language":  "fr-CI,fr;q=0.9,en-GB;q=0.8,en;q=0.7,fr-FR;q=0.6,en-US;q=0.5",
		"Cache-Control":    "max-age=0",
		"f-None-Match":     "107tr9n2kobde93",
		"Priority":         "u=0, i",
		"Sec-Ch-Ua":        "\"Not/A)Brand\";v=\"8\", \"Chromium\";v=\"126\", \"Google Chrome\";v=\"126\"",
		"Sec-Ch-Ua-Mobile": "?1", "Sec-Ch-Ua-Platform": "Android",
		"Sec-Fetch-Dest":            "document",
		"Sec-Fetch-Mode":            "navigate",
		"Sec-Fetch-Site":            "same-origin",
		"Sec-Fetch-User":            "?1",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Mobile Safari/537.36",
	}
}
