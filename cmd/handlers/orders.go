package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tgmendes/go4dummies/internal/order"
)

type OrderHandler struct {
	OrdersAPIKey string
	OrdersHost   string
}

func (oh OrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	placeOrder(w, r)
}

func HandleNewOrder(w http.ResponseWriter, r *http.Request) {
	placeOrder(w, r)
}

func placeOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("invalid method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newOrder := &order.NewOrder{}
	if err = json.Unmarshal(b, newOrder); err != nil {
		log.Println("error unmarshaling: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	httpClient := http.DefaultClient
	oc := &order.Client{
		APIKey: "some-key",
		HTTP:   httpClient,
		Host:   "http://example.com",
	}

	order, err := oc.PlaceOrder(*newOrder)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	finalOrder, err := json.Marshal(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(finalOrder)
}
