package handlers

import (
	"net/http"

	"github.com/tgmendes/go4dummies/internal/platform/db"
	"github.com/tgmendes/go4dummies/internal/platform/yelp"
)

// CrawlYelp will invoke the yelp crawler.
func CrawlYelp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	db, _ := db.New("127.0.0.1:27017/yelp", 0)

	cr := yelp.Crawler{
		BaseURL:   "https://www.yelp.com/search?find_loc=%s&start=%d",
		DB:        db,
		Locations: []string{"lisbon", "porto", "faro", "coimbra"},
		Pages:     []int{1, 10},
	}
	go cr.CollectRestaurants() // HL

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("crawling"))
}
