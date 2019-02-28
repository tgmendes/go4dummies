package order

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var ErrPlaceOrder = errors.New("Unexpected status code from order placement")

type Client struct {
	Host   string
	APIKey string
	HTTP   *http.Client
}

// FIRST OMIT
func (cl Client) PlaceOrder(newOrder NewOrder) (*Order, error) {
	b, err := json.Marshal(newOrder)
	if err != nil {
		return nil, err
	}
	// ENDFIRST OMIT

	// SEC OMIT
	req, err := http.NewRequest(http.MethodPost, cl.Host+"/publicapi/v1/send-order",
		bytes.NewBuffer(b))

	req.Header.Set("X-Access-Token", cl.APIKey)
	req.Header.Set("Content-Type", "application/json")

	res, err := cl.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	// ENDSEC OMIT

	// THIRD OMIT
	if res.StatusCode != http.StatusOK {
		return nil, ErrPlaceOrder
	}
	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	order := &Order{}
	if json.Unmarshal(resp, order); err != nil {
		return nil, err
	}
	return order, nil
}

// ENDTHIRD OMIT
