package order

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	Host string
	HTTP *http.Client
}

func (cl *Client) PlaceOrder(newOrder *NewOrder) (*Order, error) {
	b, err := json.Marshal(newOrder)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, cl.Host+"/order", bytes.NewBuffer(b))
	cl.HTTP.Do(req)
	return nil, nil
}
