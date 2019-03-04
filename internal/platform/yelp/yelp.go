package yelp

import (
	"gopkg.in/mgo.v2"
)

// DBConnection represents an abstraction for the database.
type DBConnection interface {
	Execute(collName string, f func(*mgo.Collection) error) error
}

const restaurantsCollection = "restaurants"
