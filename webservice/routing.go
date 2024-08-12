package webservice

import (
	"net/http"
)

type Routes = map[string]func(w http.ResponseWriter, r *http.Request)

type Route struct {
	Path    string
	Method  string
	Handler RouteHandler
}

func NewRoute(method string,
	path string,
	handler RouteHandler) Route {
	return Route{
		Path:    path,
		Method:  method,
		Handler: handler,
	}
}
