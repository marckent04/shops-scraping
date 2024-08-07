package cart

import "shops-scraping/webservice"

var Routes = webservice.Routes{
	"/cart":             getCart,
	"/cart/add":         addToCart,
	"/cart/delete-line": deleteCartLine,
	"/cart/clear":       clearCart,
}
