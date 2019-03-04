package yelp

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/tgmendes/go4dummies/internal/platform/db"
)

const restaurantsCollection = "restaurants"

type Restaurants []string

func (r *Restaurants) append(t string) {
	*r = append(*r, t)
}

func CollectRestaurants(locations []string, pages []int, dbConn *db.DB) {
	chRest := make(chan Restaurant)
	chDone := make(chan bool)

	for _, l := range locations {
		for _, p := range pages {
			go ConcurrentCrawl(l, p, chRest, chDone)
		}
	}

	for c := 0; c < len(locations)*len(pages); {
		select {
		case rest := <-chRest:
			rest.ID = bson.NewObjectId()
			f := func(collection *mgo.Collection) error {
				log.Println("adding record")
				return collection.Insert(&rest)
			}
			if err := dbConn.Execute(restaurantsCollection, f); err != nil {
				log.Println("insert restaurants error: ", err.Error())
			}
		case <-chDone:
			c++
		}
	}

}
