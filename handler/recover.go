package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/rafael84/shortener/service"
)

func Recover(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	alias := strings.TrimPrefix(r.URL.Path, "/a/")
	url, err := service.Recover(alias)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, `{"err":"Not Found"}`)
		accessLog(r, 404)
	} else {
		w.Header().Set("Location", url)
		w.WriteHeader(302)
		accessLog(r, 302, url)
	}
}
