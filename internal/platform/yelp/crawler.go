package yelp

import (
	"fmt"
	"net/http"
	"regexp"

	"golang.org/x/net/html"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Crawler represents a crawler for restaurants.
type Crawler struct {
	BaseURL   string
	DB        DBConnection
	Locations []string
	Pages     []int
}

// CollectRestaurants invokes the concurrent crawler to add restaurants to a given database.
// START OMIT
func (cr *Crawler) CollectRestaurants() {
	chRest := make(chan Restaurant)
	chDone := make(chan bool)

	for _, l := range cr.Locations {
		for _, p := range cr.Pages {
			go Crawl(cr.BaseURL, l, p, chRest, chDone) // HL
		}
	}

	for c := 0; c < len(cr.Locations)*len(cr.Pages); { // HL
		select { // HL
		case rest := <-chRest:
			rest.ID = bson.NewObjectId()
			f := func(collection *mgo.Collection) error {
				return collection.Insert(&rest)
			}
			cr.DB.Execute(restaurantsCollection, f) // HL
		case <-chDone:
			c++
		}
	}
}

// END OMIT

// Crawl crawls a list of restaurants for a given URL, location and start page.
// It will publish the results to a channel.
func Crawl(URL string, location string, start int, ch chan Restaurant, done chan bool) {
	url := fmt.Sprintf(URL, location, start)
	resp, _ := http.Get(url)

	defer func() {
		done <- true // HL
	}()

	b := resp.Body
	defer b.Close()

	doc, _ := html.Parse(b)

	extractRestaurant(location, start, doc, ch)
}

func extractRestaurant(location string, page int, node *html.Node, ch chan Restaurant) { // HL
	var retaurantClass = regexp.MustCompile(`heading`)

	if node.Type == html.ElementNode && node.Data == "h3" {
		for _, attr := range node.Attr {
			if attr.Key == "class" && retaurantClass.MatchString(attr.Val) {
				for c := node.FirstChild; c != nil; c = c.NextSibling {
					if c.Data == "a" {
						ch <- Restaurant{ // HL
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
