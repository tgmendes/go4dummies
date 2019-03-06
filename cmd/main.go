package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tgmendes/go4dummies/cmd/handlers"
)

func main() {
	APIKey := os.Getenv("API_KEY")
	Host := os.Getenv("HOST")

	oh := handlers.OrderHandler{
		EatStreetAPIKey: APIKey,
		EatStreetHost:   Host,
	}

	log.Println("starting server")
	http.Handle("/api/orders", oh)                                // HL
	http.HandleFunc("/api/restaurants/crawl", handlers.CrawlYelp) // HL
	http.ListenAndServe(":8080", nil)                             // HL
}
