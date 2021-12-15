package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/product", handler.ProductIdHandler)

	log.Println("Web listening on port: 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
