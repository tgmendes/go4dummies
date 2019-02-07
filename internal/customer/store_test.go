package customer_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/tgmendes/go4dummies/internal/customer"

	"github.com/tgmendes/go4dummies/internal/mocks"
)

var mockDb *mocks.DB
var expCustomers []customer.Customer

func init() {
	expCustomers := []customer.Customer{
		customer.Customer{
			ID:    1,
			Name:  "Joey Banana",
			Email: "joey@banana.com",
		},
		customer.Customer{
			ID:    1,
			Name:  "Jane Potato",
			Email: "jane@potato.com",
		},
	}
	mockDb = &mocks.DB{
		AddFn: func(entity []byte) ([]byte, error) {
			newCust := &customer.NewCustomer{}
			json.Unmarshal(entity, newCust)
			cust := &customer.Customer{
				ID:    1,
				Name:  newCust.Name,
				Email: newCust.Email,
			}
			return json.Marshal(cust)
		},
		ListFn: func() ([]byte, error) {
			return json.Marshal(expCustomers)
		},
	}
}

func TestList(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "list retrieves the list of customers",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			s := customer.NewStore(mockDb)
			cust, err := s.List()

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			fmt.Println(())
			if !reflect.DeepEqual(cust, expCustomers) {
				t.Logf("want: %+v", expCustomers)
				t.Logf("got: %+v", cust)
				t.Fatal("lists do not match")
			}

		})
	}
}
