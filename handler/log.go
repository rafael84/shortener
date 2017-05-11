package handler

import (
	"log"
	"net/http"
)

func accessLog(r *http.Request, status int, extra ...interface{}) {
	log.Printf("method[%v] uri[%v] remoteAddr[%v] userAgent[%v] status[%v] extra%v",
		r.Method, r.RequestURI, r.RemoteAddr, r.UserAgent(), status, extra,
	)
}
