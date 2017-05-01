package handler

import (
	"net/http"
	"strings"

	"github.com/rafael84/shortener/service"
)

func Recover(w http.ResponseWriter, r *http.Request) {
	alias := strings.TrimPrefix(r.URL.Path, "/a/")
	url, err := service.Recover(alias)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
	} else {
		w.Header().Set("Location", url)
		w.WriteHeader(302)
	}
}
