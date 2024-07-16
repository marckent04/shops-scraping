package ZARA

import "shops-scraping/scraping/common"

const searchUrl = "https://www.zara.com/fr/fr/search?searchTerm=%s&section=%s"
const articleSelector = "li.product-grid-product"
const shopName = "ZARA"

var genders = map[common.Gender]string{
	common.MAN:   "MAN",
	common.WOMAN: "WOMAN",
}
