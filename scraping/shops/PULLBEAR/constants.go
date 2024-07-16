package PULLBEAR

import "shops-scraping/scraping/common"

const searchUrl = "https://www.pullandbear.com/fr/%s?q=%s"
const articleSelector = "grid-product"
const shopName = "PULL&BEAR"

var genders = map[common.Gender]string{
	common.MAN:   "homme-n6228",
	common.WOMAN: "femme-n6417",
}
