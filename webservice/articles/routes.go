package articlesController

import "shops-scraping/webservice"

var Routes = webservice.Routes{
	"/articles":        searchByKeywords,
	"/{shop}/articles": searchByShop,
}
