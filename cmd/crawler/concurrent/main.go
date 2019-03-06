package main

import (
	"github.com/tgmendes/go4dummies/internal/platform/yelp"

	"github.com/tgmendes/go4dummies/internal/platform/db"
)

func main() {
	db, _ := db.New("127.0.0.1:27017/yelp", 0)

	cr := yelp.Crawler{
		BaseURL: "https://www.yelp.com/search?find_loc=%s&start=%d",
		DB:      db,
		Locations: []string{"lisbon",
			"porto",
			"faro",
			"coimbra",
		},
		Pages: []int{
			1,
			10,
		},
	}
	cr.CollectRestaurants()
}
