package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tgmendes/go4dummies/internal/order"
)

// ORDERONE OMIT
type OrderHandler struct {
	EatStreetAPIKey string
	EatStreetHost   string
}

func (oh OrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { // HL
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	b, _ := ioutil.ReadAll(r.Body)

	newOrder := &order.NewOrder{}
	if err := json.Unmarshal(b, newOrder); err != nil { // HL
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// ENDORDERONE OMIT
	// ORDERTWO OMIT

	oc := &order.Client{
		APIKey: oh.EatStreetAPIKey,
		HTTP:   http.DefaultClient,
		Host:   oh.EatStreetHost,
	}

	order, err := oc.PlaceOrder(*newOrder) // HL
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	finalOrder, _ := json.Marshal(order) // HL

	w.Write(finalOrder)

}

// ENDORDERTWO OMIT
