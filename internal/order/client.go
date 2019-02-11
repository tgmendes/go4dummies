package order

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var ErrPlaceOrder = errors.New("Unexpected status code from order placement")

type Client struct {
	Host   string
	HTTP   *http.Client
	APIKey string
}

func (cl Client) PlaceOrder(newOrder NewOrder) (*Order, error) {
	b, err := json.Marshal(newOrder)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, cl.Host+"/publicapi/v1/send-order", bytes.NewBuffer(b))
	req.Header.Set("X-Access-Token", cl.APIKey)
	req.Header.Set("Content-Type", "application/json")

	res, err := cl.HTTP.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, ErrPlaceOrder
	}
	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(res.Body)

	order := &Order{}
	if json.Unmarshal(resp, order); err != nil {
		return nil, err
	}
	return order, nil
}
