package yelp

import (
	"fmt"
	"net/http"
	"regexp"

	"golang.org/x/net/html"
)

func ConcurrentCrawl(location string, start int, ch chan Restaurant, done chan bool) {
	url := fmt.Sprintf("https://www.yelp.com/search?find_loc=%s&start=%d", location, start)

	resp, _ := http.Get(url)

	defer func() {
		done <- true
	}()

	b := resp.Body
	defer b.Close()

	doc, _ := html.Parse(b)

	extractRestaurant(location, start, doc, ch)
}

func extractRestaurant(location string, page int, node *html.Node, ch chan Restaurant) {
	var retaurantClass = regexp.MustCompile(`heading`)

	if node.Type == html.ElementNode && node.Data == "h3" {
		for _, attr := range node.Attr {
			if attr.Key == "class" && retaurantClass.MatchString(attr.Val) {
				for c := node.FirstChild; c != nil; c = c.NextSibling {
					if c.Data == "a" {
						ch <- Restaurant{
							Name:     c.FirstChild.Data,
							Location: location,
							Page:     page,
						}

					}
				}
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		extractRestaurant(location, page, c, ch)
	}
}
