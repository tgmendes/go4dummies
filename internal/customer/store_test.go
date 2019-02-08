package customer_test

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	. "github.com/tgmendes/go4dummies/internal/customer"

	"github.com/tgmendes/go4dummies/internal/mocks"
)

var expCustomers []Customer

func init() {
	expCustomers = []Customer{
		Customer{
			ID:    1,
			Name:  "Joey Banana",
			Email: "joey@banana.com",
		},
		Customer{
			ID:    1,
			Name:  "Jane Potato",
			Email: "jane@potato.com",
		},
	}

}

func TestList(t *testing.T) {
	testCases := []struct {
		desc   string
		listFn func() ([]byte, error)
		expErr bool
	}{
		{
			desc: "retrieves the list of customers",
			listFn: func() ([]byte, error) {
				return json.Marshal(expCustomers)
			},
		},
		{
			desc: "returns an error when DB errors",
			listFn: func() ([]byte, error) {
				return nil, errors.New("an error ocurred")
			},
			expErr: true,
		},
		{
			desc: "returns an error when unmarshalling bad entity",
			listFn: func() ([]byte, error) {
				return []byte("potato"), nil
			},
			expErr: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// given
			mockDb := &mocks.DB{
				ListFn: tC.listFn,
			}

			s := &Store{
				DB: mockDb,
			}

			// when
			cust, err := s.List()

			// then
			if tC.expErr {
				if err == nil {
					t.Fatalf("expected an error")
				}

			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				if !reflect.DeepEqual(cust, expCustomers) {
					t.Logf("want: %+v", expCustomers)
					t.Logf("got: %+v", cust)
					t.Fatal("lists do not match")
				}
			}
			if mockDb.ListFnCalled != 1 {
				t.Fatalf("list should have been called 1 time")
			}

		})
	}
}

func TestCreate(t *testing.T) {
	dbCust := &Customer{
		ID:    1,
		Name:  "joe",
		Email: "joe@banana.com",
	}
	newCust := &NewCustomer{
		Name:  "joe",
		Email: "joe@banana.com",
	}

	testCases := []struct {
		desc   string
		expErr bool
		addFn  func(entity []byte) ([]byte, error)
	}{
		{
			desc: "creates successfully a customer and returns the new one",
			addFn: func(entity []byte) ([]byte, error) {
				newCust := &NewCustomer{}
				json.Unmarshal(entity, newCust)

				return json.Marshal(dbCust)
			},
		},
		{
			desc: "returns an error when DB errors",
			addFn: func(entity []byte) ([]byte, error) {
				return nil, errors.New("an error ocurred")
			},
			expErr: true,
		},
		{
			desc: "returns an error when unmarshalling bad entity",
			addFn: func(entity []byte) ([]byte, error) {
				return []byte("potato"), nil
			},
			expErr: true,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			mockDb := &mocks.DB{
				AddFn: tC.addFn,
			}

			s := &Store{
				DB: mockDb,
			}

			// when
			cust, err := s.Create(newCust)

			// then
			if tC.expErr {
				if err == nil {
					t.Fatalf("expected an error")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				if !reflect.DeepEqual(cust, dbCust) {
					t.Logf("want: %+v", dbCust)
					t.Logf("got: %+v", cust)
					t.Fatal("lists do not match")
				}
			}
			if mockDb.AddFnCalled != 1 {
				t.Fatalf("list should have been called 1 time")
			}

		})
	}
}
