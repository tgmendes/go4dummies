package main

import (
	"fmt"

	"github.com/tgmendes/go4dummies/internal/platform/yelp"
)

func main() {
	allRests := []string{}
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
	for _, l := range locations {
		for _, p := range pages {

			rests := yelp.SimpleCrawl(l, p)
			for _, r := range *rests {
				allRests = append(allRests, r)
			}
		}
	}
	fmt.Println(allRests)
}
