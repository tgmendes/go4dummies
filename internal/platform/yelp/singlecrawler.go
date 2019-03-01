package yelp

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"golang.org/x/net/html"
)

func SimpleCrawl(location string, start int) *Restaurants {
	url := fmt.Sprintf("https://www.yelp.com/search?find_loc=%s&start=%d", location, start)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	doc, _ := html.Parse(resp.Body)

	rest := &Restaurants{}
	extractRestaurants(doc, rest)
	return rest
}

func extractRestaurants(node *html.Node, rest *Restaurants) {
	var retaurantClass = regexp.MustCompile(`heading`)

	if node.Type == html.ElementNode && node.Data == "h3" {
		for _, attr := range node.Attr {
			if attr.Key == "class" && retaurantClass.MatchString(attr.Val) {
				for c := node.FirstChild; c != nil; c = c.NextSibling {
					if c.Data == "a" {
						rest.append(c.FirstChild.Data)
					}
				}
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		extractRestaurants(c, rest)
	}
}
