package customer

import (
	"encoding/json"
	"log"
)

// Store represents a customer service for handling customer data.
type Store struct {
	DB interface {
		Add(entity []byte) ([]byte, error)
		List() ([]byte, error)
		View(id int) ([]byte, error)
	}
}

// List returns a list of all customers in the database.
func (s *Store) List() ([]Customer, error) {
	cust := []Customer{}
	res, err := s.DB.List()
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
	b, err := json.Marshal(newCust)
	if err != nil {
		return nil, err
	}

	b, err = s.DB.Add(b)
	if err != nil {
		return nil, err
	}

	cust := &Customer{}
	if err := json.Unmarshal(b, cust); err != nil {
		return nil, err
	}
	return cust, nil
}
