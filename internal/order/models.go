package order

import (
	"time"

	"github.com/tgmendes/go4dummies/internal/customer"
)

type Order struct {
	ID       int                `json:"id"`
	Customer *customer.Customer `json:"customer"`
	Menu     string             `json:"menu"`
	Ordered  time.Time          `json:"ordered"`
}

type NewOrder struct {
	Customer *customer.Customer `json:"customer"`
	Menu     string             `json:"menu"`
}
