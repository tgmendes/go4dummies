package main

import (
	"log"
	"net/http"

	"github.com/tgmendes/go4dummies/cmd/handlers"
)

func main() {
	log.Println("starting server")
	http.HandleFunc("/api/orders", handlers.HandleNewOrder)
	http.ListenAndServe(":8080", nil)
}
