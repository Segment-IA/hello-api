package main

import (
	"log"
	"net/http"
	"github.com/segment-ia/hello-api/handler"
)

func main() {
	http.HandleFunc("/", handler.HelloHandler)

	log.Println("🚀 Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
