package db

import (
	"github.com/boltdb/bolt"
)

type DB struct {
	database *bolt.DB
}
