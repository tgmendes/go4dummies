package yelp

import (
	"fmt"
	"net/http"
	"regexp"

	"golang.org/x/net/html"
)

// Restaurants represents a list of restaurants from yelp.
type Restaurants []string

func (r *Restaurants) append(t string) {
	*r = append(*r, t)
}

// SimpleCrawl will crawl the yelp page without concurrency.
func SimpleCrawl(location string, start int) *Restaurants {
	url := fmt.Sprintf("https://www.yelp.com/search?find_loc=%s&start=%d", location, start)
	resp, _ := http.Get(url)

	doc, _ := html.Parse(resp.Body)

	rest := &Restaurants{}
	// this will parse the HTML document for restaurant names
	// and append results to a restaurant list.
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
