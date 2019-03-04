package main

import (
	"log"

	"github.com/tgmendes/go4dummies/internal/platform/yelp"

	"github.com/tgmendes/go4dummies/internal/platform/db"
)

func main() {
	db, err := db.New("127.0.0.1:27017/yelp", 0)

	if err != nil {
		log.Fatal(err)
	}

	locations := []string{
		"lisbon",
		"porto",
		"faro",
		"coimbra",
	}

	pages := []int{
		1,
		10,
	}

	yelp.CollectRestaurants(locations, pages, db)
}
