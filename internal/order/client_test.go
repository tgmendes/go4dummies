package order_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	. "github.com/tgmendes/go4dummies/internal/order"
)

// Success and failure markers.
const (
	Success = "\u2713"
	Failed  = "\u2717"
)

var testOrder *Order
var rawTestOrder []byte
var newOrder NewOrder

func init() {
	testOrder = &Order{}
	rawTestOrder, _ = ioutil.ReadFile("test-fixtures/test-order.json")
	json.Unmarshal(rawTestOrder, testOrder)

	newOrder = NewOrder{
		Address: APIKey{
			APIKey: "someAddressKey",
		},
		Card: APIKey{
			APIKey: "someCardKey",
		},
		Items:   testOrder.Items,
		Method:  testOrder.Method,
		Payment: testOrder.Payment,
		Recipient: APIKey{
			APIKey: testOrder.RecipientAPIKey,
		},
		RestaurantAPIKey: testOrder.RestaurantAPIKey,
		Test:             true,
	}
}

func TestPlaceOrder(t *testing.T) {
	// expected
	expAuth := "someApiKey"
	expContent := "application/json"
	expPath := "/publicapi/v1/send-order"
	// actual
	var actAuth string
	var actContent string
	var actPath string
	var actReqBody []byte

	t.Log("Given an order mock server and request.")
	{

		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			actPath = r.URL.Path
			actAuth = r.Header.Get("X-Access-Token")
			actContent = r.Header.Get("Content-Type")
			actReqBody, _ = ioutil.ReadAll(r.Body)
			w.Write(rawTestOrder)
		}))
		defer srv.Close()

		httpCl := http.DefaultClient
		cl := &Client{
			Host:   srv.URL,
			HTTP:   httpCl,
			APIKey: expAuth,
		}

		t.Log("\tWhen placing an order.")
		{
			res, err := cl.PlaceOrder(newOrder)

			if err != nil {
				t.Fatalf("\t%s\tShould be able to place an order : %s.", Failed, err)
			} else {

				t.Logf("\t%s\tShould be able to place an order.", Success)
			}

			if actPath != expPath {
				t.Log("\t\tGot : ", actPath)
				t.Log("\t\tWant: ", expPath)
				t.Errorf("\t%s\tShould have correct auth header.", Failed)
			} else {

				t.Logf("\t%s\tShould have correct auth header.", Success)
			}

			if actAuth != expAuth {
				t.Log("\t\tGot : ", actAuth)
				t.Log("\t\tWant: ", expAuth)
				t.Errorf("\t%s\tShould have correct auth header.", Failed)
			} else {

				t.Logf("\t%s\tShould have correct auth header.", Success)
			}

			if actContent != expContent {
				t.Log("\t\tGot : ", actContent)
				t.Log("\t\tWant: ", expContent)
				t.Errorf("\t%s\tShould have correct content type.", Failed)
			} else {

				t.Logf("\t%s\tShould have correct content type.", Success)
			}

			b, _ := json.Marshal(newOrder)
			if !reflect.DeepEqual(actReqBody, b) {
				t.Logf("\t\tGot : %+v", actReqBody)
				t.Logf("\t\tWant: %+v", b)
				t.Errorf("\t%s\tShould have matching orders.", Failed)
			} else {

				t.Errorf("\t%s\tShould have matching orders.", Success)
			}

			if !reflect.DeepEqual(res, testOrder) {
				t.Logf("\t\tGot : %+v", res)
				t.Logf("\t\tWant: %+v", testOrder)
				t.Errorf("\t%s\tShould have matching orders.", Failed)
			} else {

				t.Logf("\t%s\tShould have matching orders.", Success)
			}

		}
	}
}

func Test(t *testing.T) {
	testCases := []struct {
		desc       string
		statusCode int
		expErr     error
	}{
		{
			desc:       "should error for bad status code json",
			statusCode: http.StatusInternalServerError,
			expErr:     ErrPlaceOrder,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Log("Given an order mock server and request.")
			{
				srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(tC.statusCode)
				}))
				defer srv.Close()

				httpCl := http.DefaultClient
				cl := &Client{
					Host:   srv.URL,
					HTTP:   httpCl,
					APIKey: "banana",
				}
				t.Log("\tWhen placing an order.")
				{
					_, err := cl.PlaceOrder(newOrder)
					if err != tC.expErr {
						t.Logf("\t\tGot : %+v", err)
						t.Logf("\t\tWant: %+v", tC.expErr)
						t.Errorf("\t%s\tShould have matching errors.", Failed)
					} else {
						t.Logf("\t%s\tShould have matching errors.", Success)
					}
				}
			}
		})
	}

}

func TestPlaceOrderErrors(t *testing.T) {
	newOrder := NewOrder{
		Address: APIKey{
			APIKey: "someAddressKey",
		},
		Card: APIKey{
			APIKey: "someCardKey",
		},
		Items:   testOrder.Items,
		Method:  testOrder.Method,
		Payment: testOrder.Payment,
		Recipient: APIKey{
			APIKey: testOrder.RecipientAPIKey,
		},
		RestaurantAPIKey: testOrder.RestaurantAPIKey,
		Test:             true,
	}

	// expected
	expAuth := "someApiKey"
	expContent := "application/json"
	expPath := "/publicapi/v1/send-order"
	// actual
	var actAuth string
	var actContent string
	var actPath string
	var actReqBody []byte

	t.Log("Given an order mock server and request.")
	{

		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			actPath = r.URL.Path
			actAuth = r.Header.Get("X-Access-Token")
			actContent = r.Header.Get("Content-Type")
			actReqBody, _ = ioutil.ReadAll(r.Body)
			w.Write(rawTestOrder)
		}))
		defer srv.Close()

		httpCl := http.DefaultClient
		cl := &Client{
			Host:   srv.URL,
			HTTP:   httpCl,
			APIKey: expAuth,
		}

		t.Log("\tWhen placing an order.")
		{
			res, err := cl.PlaceOrder(newOrder)

			if err != nil {
				t.Fatalf("\t%s\tShould be able to place an order : %s.", Failed, err)
			}
			t.Logf("\t%s\tShould be able to place an order.", Success)

			if actPath != expPath {
				t.Log("\t\tGot : ", actPath)
				t.Log("\t\tWant: ", expPath)
				t.Errorf("\t%s\tShould have correct auth header.", Failed)
			}
			t.Logf("\t%s\tShould have correct auth header.", Success)

			if actAuth != expAuth {
				t.Log("\t\tGot : ", actAuth)
				t.Log("\t\tWant: ", expAuth)
				t.Errorf("\t%s\tShould have correct auth header.", Failed)
			}
			t.Logf("\t%s\tShould have correct auth header.", Success)

			if actContent != expContent {
				t.Log("\t\tGot : ", actContent)
				t.Log("\t\tWant: ", expContent)
				t.Errorf("\t%s\tShould have correct content type.", Failed)
			}
			t.Logf("\t%s\tShould have correct content type.", Success)

			b, _ := json.Marshal(newOrder)
			if !reflect.DeepEqual(actReqBody, b) {
				t.Logf("\t\tGot : %+v", actReqBody)
				t.Logf("\t\tWant: %+v", b)
				t.Errorf("\t%s\tShould have matching orders.", Failed)
			}

			if !reflect.DeepEqual(res, testOrder) {
				t.Logf("\t\tGot : %+v", res)
				t.Logf("\t\tWant: %+v", testOrder)
				t.Errorf("\t%s\tShould have matching orders.", Failed)
			}
			t.Logf("\t%s\tShould have matching orders.", Failed)

		}
	}
}
