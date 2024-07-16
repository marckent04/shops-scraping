package HM

import "shops-scraping/scraping/common"

const productsListSelector = "#main-content > section > ul"
const searchUrl = "https://www2.hm.com/fr_fr/search-results.html"
const productSelector = "article[data-articlecode]"
const shopName = "H&M"

var genders = map[common.Gender]string{
	common.MAN:   "men_all",
	common.WOMAN: "ladies_all",
}
