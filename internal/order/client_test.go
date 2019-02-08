package order_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tgmendes/go4dummies/internal/customer"
	. "github.com/tgmendes/go4dummies/internal/order"
)

func TestOrder(t *testing.T) {

	testCust := &customer.Customer{
		ID:    1,
		Name:  "joe pher",
		Email: "joepher@gopherinos.com",
	}
	testMenu := "fish and chips"
	testCases := []struct {
		desc string
	}{
		{
			desc: "places order successfully",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// given
			// expected
			expOrder := `{
				"id": 1,
				"customer": {
					"id": 1,
					"name": "joe pher",
					"email": "joepher@gopherinos.com"
				},
				"menu": "fish and chips"
			}`
			expMeth := http.MethodPost
			// expPath := "/orders"

			var actNewOrder *NewOrder
			var actMeth string
			// var actPath string
			// var actAuth string

			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				actMeth = r.Method
				// actPath = r.URL.Path
				// actAuth = r.Header.Get("Authorization")
				// we are ignoring the second value
				b, _ := ioutil.ReadAll(r.Body)

				actNewOrder = &NewOrder{}
				if err := json.Unmarshal(b, actNewOrder); err != nil {
					t.Fatalf("unexpected error %v", err)
				}

				w.Write([]byte(expOrder))

			}))
			defer srv.Close()

			cl := &Client{
				Host: srv.URL,
				HTTP: http.DefaultClient,
			}

			// when
			cl.PlaceOrder(&NewOrder{
				Customer: testCust,
				Menu:     testMenu,
			})

			// then
			if expMeth != actMeth {
				t.Logf("want: %+v", expMeth)
				t.Logf("got: %+v", actMeth)
				t.Fatal("methods do not match")
			}

		})
	}
}
