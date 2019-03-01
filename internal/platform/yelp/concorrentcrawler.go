package yelp

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"golang.org/x/net/html"
)

func ConcorrentCrawl(location string, start int, ch chan string, done chan bool) {
	url := fmt.Sprintf("https://www.yelp.com/search?find_loc=%s&start=%d", location, start)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		done <- true
	}()

	b := resp.Body
	defer b.Close()

	doc, _ := html.Parse(b)

	rest := &Restaurants{}
	extractRestaurant(doc, rest, ch)
}

func extractRestaurant(node *html.Node, rest *Restaurants, ch chan string) {
	var retaurantClass = regexp.MustCompile(`heading`)

	if node.Type == html.ElementNode && node.Data == "h3" {
		for _, attr := range node.Attr {
			if attr.Key == "class" && retaurantClass.MatchString(attr.Val) {
				for c := node.FirstChild; c != nil; c = c.NextSibling {
					if c.Data == "a" {
						ch <- c.FirstChild.Data
					}
				}
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		extractRestaurant(c, rest, ch)
	}
}
