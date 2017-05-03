package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/rafael84/shortener/handler"
	"github.com/rafael84/shortener/persistence"
	"github.com/rafael84/shortener/service"
)

var (
	redisAddr = flag.String("redis-addr", "", "host:port of a running redis server")
)

func main() {
	flag.Parse()

	if *redisAddr != "" {
		service.Storage = persistence.NewRedis(*redisAddr, "", 0)
	}

	http.HandleFunc("/create", handler.Create)
	http.HandleFunc("/a/", handler.Recover)
	log.Println("server running at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
