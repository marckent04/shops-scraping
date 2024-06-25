package webservice

import (
	"fmt"
	"net/http"
	"strings"
)

type Routes = map[string]func(w http.ResponseWriter, r *http.Request)

func RegisterApiRoutes(routes Routes) {
	apiPrefix := "/api"
	for url, handler := range routes {
		url, _ = strings.CutPrefix(url, "/")
		http.HandleFunc(fmt.Sprintf("%s/%s", apiPrefix, url), handler)
	}
}
