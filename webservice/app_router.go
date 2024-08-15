package webservice

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"shops-scraping/shared"
	"slices"
)

type RouteHandler = http.HandlerFunc

type HTTPRouter struct {
	prefix string
	routes []Route
}

func (r *HTTPRouter) Get(url string, handler RouteHandler) {
	r.registerHandler(http.MethodGet, url, handler)
}

func (r *HTTPRouter) Put(url string, handler RouteHandler) {
	r.registerHandler(http.MethodPut, url, handler)
}

func (r *HTTPRouter) Post(url string, handler RouteHandler) {
	r.registerHandler(http.MethodPost, url, handler)
}

func (r *HTTPRouter) SetGlobalPrefix(prefix string) {
	r.prefix = prefix
}

func (r *HTTPRouter) Delete(url string, handler RouteHandler) {
	r.registerHandler(http.MethodDelete, url, handler)
}

func (r *HTTPRouter) registerHandler(method string, url string, handler RouteHandler) {
	alreadyRegistered := slices.ContainsFunc(r.routes, func(route Route) bool {
		return route.Path == url && route.Method == method
	})

	if alreadyRegistered {
		log.Panicf("handler for %s with %s method already registered", url, method)
	}

	r.routes = append(r.routes, newRoute(method, url, handler))
}

func (r *HTTPRouter) compile() {
	var paths []string
	for _, route := range r.routes {
		if !slices.Contains(paths, route.Path) {
			paths = append(paths, route.Path)
		}
	}

	for _, path := range paths {
		http.HandleFunc(fmt.Sprintf("%s%s", r.prefix, path), func(w http.ResponseWriter, req *http.Request) {
			routes := shared.SlicesFilter(r.routes, func(r Route) bool {
				return r.Path == path
			})

			for _, route := range routes {
				if req.Method == route.Method {
					route.Handler(w, req)
				}
			}
		})
	}

}

func (r *HTTPRouter) Listen(port string) {
	r.compile()
	log.Println("server is launching on port", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}

func newHttpRouter() HTTPRouter {
	return HTTPRouter{}
}
