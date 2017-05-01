package main

import (
	"log"
	"net/http"

	"github.com/rafael84/shortener/handler"
)

func main() {
	http.HandleFunc("/create", handler.Create)
	http.HandleFunc("/a/", handler.Recover)
	log.Println("server running at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
