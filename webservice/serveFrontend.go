package webservice

import (
	"net/http"
)

func serveFrontend() {
	fs := http.FileServer(http.Dir("front/dist"))
	http.Handle("/assets/", fs)
	routes := []string{"/", "/search"}

	for _, route := range routes {
		http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "front/dist/index.html")
		})
	}
}
