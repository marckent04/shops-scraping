package webservice

import (
	"net/http"
)

type Routes = map[string]func(w http.ResponseWriter, r *http.Request)
type RouteHandler = func(http.ResponseWriter, *http.Request) httpResponse

type Route struct {
	Path    string
	Method  string
	Handler RouteHandler
}

func newRoute(method string,
	path string,
	handler RouteHandler) Route {
	return Route{
		Path:    path,
		Method:  method,
		Handler: handler,
	}
}

type httpResponse struct {
	Data   interface{}
	Status int
}
