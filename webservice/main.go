package webservice

import (
	"os"
)

func Start() {
	serveFrontend()

	port := os.Getenv("PORT")
	app := newHttpRouter()

	app.SetGlobalPrefix("/api")
	app.Get("/articles", searchByShops)
	app.Get("/cart", getCart)
	app.Delete("/cart/line", deleteCartLine)
	app.Get("/shops", getEnabledShops)
	app.Delete("/cart", clearCart)
	app.Post("/cart", addToCart)

	app.Listen(port)
}
