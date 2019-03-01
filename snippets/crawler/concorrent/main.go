package main

import (
	"fmt"

	"github.com/tgmendes/go4dummies/internal/platform/yelp"
)

func main() {
	// allRests := []string{}
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
	chRest := make(chan string)
	chDone := make(chan bool)
	for _, l := range locations {
		for _, p := range pages {
			go yelp.ConcorrentCrawl(l, p, chRest, chDone)
		}
	}
	for c := 0; c < len(locations)*len(pages); {
		select {
		case rest := <-chRest:
			fmt.Println(rest)
		case <-chDone:
			c++
		}
	}
}
