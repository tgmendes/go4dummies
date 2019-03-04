package yelp

import "gopkg.in/mgo.v2/bson"

// Restaurant represents a restaurant gathered from yelp.
type Restaurant struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Name     string        `bson:"name" json:"name"`
	Page     int           `bson:"page" json:"page"`
	Location string        `bson:"location" json:"location"`
}
