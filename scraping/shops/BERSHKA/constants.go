package BERSHKA

import "shops-scraping/scraping/common"

const searchUrl = "https://www.bershka.com/fr/q"
const articleSelector = ".search-product-card"
const shopName = "Bershka"

var genders = map[common.Gender]string{
	common.MAN:   "man",
	common.WOMAN: "woman",
}
