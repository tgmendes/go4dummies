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

	log.Println("starting server...")
	mux := http.NewServeMux() // HL

	mux.HandleFunc("/api/v2/restaurants/crawl", handlers.CrawlYelp) // HL
	mux.Handle("/api/v2/orders", oh)                                // HL
	http.ListenAndServe(":8080", mux)                               // HL
}
