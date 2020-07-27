package plover

import (
	"net/http"
	"net/url"
	"strings"
)

type static struct {
	PATH        string
	HttpHandler http.Handler
}

type Assets struct {
	static []static
}

func (assets *Assets) AnalysisUri(PATH string) string {
	PATH = strings.ToLower(strings.Trim(PATH, "/"))
	if strings.EqualFold(PATH, "") {
		PATH = "/"
	}
	return PATH
}

func (assets *Assets) Add(PATH string, handler http.Handler) {
	PATH = assets.AnalysisUri(PATH)
	if assets.static == nil {
		assets.static = make([]static, 0)
	}
	assets.static = append(assets.static, static{
		PATH:        PATH,
		HttpHandler: handler,
	})
}
func (assets *Assets) Handle(w http.ResponseWriter, r *http.Request) (hasStatic bool) {
	for _, static := range assets.static {
		PATH := strings.ToLower(strings.Trim(r.URL.Path, "/"))
		if strings.HasPrefix(PATH, static.PATH) {
			p := strings.TrimPrefix(strings.Trim(r.URL.Path, "/"), static.PATH)
			if len(p) < len(PATH) {
				r2 := new(http.Request)
				*r2 = *r
				r2.URL = new(url.URL)
				*r2.URL = *r.URL
				r2.URL.Path = p
				static.HttpHandler.ServeHTTP(w, r2)
				return true
			}
		}
	}
	return false
}
