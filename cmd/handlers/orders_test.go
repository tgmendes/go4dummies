package handlers_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/tgmendes/go4dummies/internal/order"

	"github.com/tgmendes/go4dummies/cmd/handlers"
)

var rawTestOrder []byte
var testOrder *order.Order
var newOrder []byte

func init() {
	testOrder = &order.Order{}
	rawTestOrder, _ = ioutil.ReadFile("../../internal/order/test-fixtures/test-order.json")
	json.Unmarshal(rawTestOrder, testOrder)

	no := order.NewOrder{
		Address: order.APIKey{
			APIKey: "someAddressKey",
		},
		Card: order.APIKey{
			APIKey: "someCardKey",
		},
		Items:   testOrder.Items,
		Method:  testOrder.Method,
		Payment: testOrder.Payment,
		Recipient: order.APIKey{
			APIKey: testOrder.RecipientAPIKey,
		},
		RestaurantAPIKey: testOrder.RestaurantAPIKey,
		Test:             true,
	}
	newOrder, _ = json.Marshal(no)
}

// TDT OMIT
func TestOrders(t *testing.T) {
	testCases := []struct {
		desc     string
		method   string
		expCode  int
		expOrder *order.Order
	}{
		{
			desc:     "happy path",
			method:   http.MethodPost,
			expCode:  http.StatusOK,
			expOrder: testOrder,
		}, {
			desc:     "wrong method",
			method:   http.MethodGet,
			expCode:  http.StatusMethodNotAllowed,
			expOrder: &order.Order{},
		},
	}
	// ENDTDT OMIT
	// RUN OMIT
	for _, tC := range testCases { // HL
		t.Run(tC.desc, func(t *testing.T) {
			// GIVEN
			// Test EatStreet Server
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write(rawTestOrder)
			}))
			defer srv.Close()

			req, _ := http.NewRequest(tC.method, "http://example.com", bytes.NewBuffer(newOrder))
			w := httptest.NewRecorder() // HL

			h := handlers.OrderHandler{
				EatStreetAPIKey: "dummy",
				EatStreetHost:   srv.URL,
			}

			// WHEN
			h.ServeHTTP(w, req) // HL

			resp, _ := ioutil.ReadAll(w.Body)
			orderResp := &order.Order{}
			json.Unmarshal(resp, orderResp)
			// ENDRUN OMIT
			// ASSERT OMIT
			// THEN
			if w.Code != tC.expCode {
				t.Logf("expected: %d", tC.expCode)
				t.Logf("got: %d", w.Code)
				t.Fatalf("response code does not match")
			}

			if !reflect.DeepEqual(tC.expOrder, orderResp) {
				t.Log(tC.expOrder)
				t.Log(orderResp)
				t.Fatalf("unexpected response")
			}
			// ENDASSERT OMIT
		})
	}
}
