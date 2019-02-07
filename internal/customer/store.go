package customer

import (
	"encoding/json"
	"log"

	"github.com/tgmendes/go4dummies/internal"
)

// Store represents a customer service for handling customer data.
type Store struct {
	db restaurant.DatabaseConnector
}

// NewStore creates a new customer Store with a given database.
func NewStore(db restaurant.DatabaseConnector) *Store {
	return &Store{
		db: db,
	}
}

// List returns a list of all customers in the database.
func (s *Store) List() ([]Customer, error) {
	cust := []Customer{}
	res, err := s.db.List()
	if err != nil {
		log.Printf("error retrieving customers: %s", err.Error())
		return nil, err
	}

	if err = json.Unmarshal(res, &cust); err != nil {
		log.Printf("error unmarshaling customers: %s", err.Error())
		return nil, err
	}
	return cust, nil
}

// Retrieve returns a customer for a given ID.
func (s *Store) Retrieve(id int) (*Customer, error) {
	return nil, nil
}

// Create creates a new customer in the db.
func (s *Store) Create(newCust *NewCustomer) (*Customer, error) {
	return nil, nil
}

// Update updates the customer with a given ID with the new provided data.
func (s *Store) Update(id int, updatedCust *NewCustomer) error {
	return nil
}

// Delete removes a customer with a given ID from the databse.
func (s *Store) Delete(id int) error {
	return nil
}
