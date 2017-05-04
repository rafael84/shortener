package main

import (
	"log"
	"net/http"

	"github.com/rafael84/shortener/config"
	"github.com/rafael84/shortener/handler"
	"github.com/rafael84/shortener/persistence"
	"github.com/rafael84/shortener/service"
)

func main() {
	if config.Data.Redis.Addr != "" {
		service.Storage = persistence.NewRedis(
			config.Data.Redis.Addr,
			config.Data.Redis.Password,
			config.Data.Redis.DB,
		)
	}

	http.HandleFunc("/create", handler.Create)
	http.HandleFunc("/a/", handler.Recover)

	log.Println("server is running", config.Data.Server.Addr)
	log.Fatal(http.ListenAndServe(config.Data.Server.Addr, nil))
}
