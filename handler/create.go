package handler

import (
	"fmt"
	"net/http"

	"github.com/rafael84/shortener/service"
)

func Create(w http.ResponseWriter, r *http.Request) {
	u := r.FormValue("url")
	a := r.FormValue("alias")

	w.Header().Set("Content-Type", "application/json")

	if a != "" {
		u, err := service.Recover(a)
		if err == nil && u != "" {
			w.WriteHeader(400)
			fmt.Fprintf(w, `{"err":"%v"}`, "alias already taken")
			return
		}
	}

	alias, err := service.Create(u, a)
	if err != nil {
		fmt.Fprintf(w, `{"err":"%v"}`, err)
	} else {
		fmt.Fprintf(w, `{"alias":"%s","url":"%s","timeTaken":"1ms"}`, alias, u)
	}
}
