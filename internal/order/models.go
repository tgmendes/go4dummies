package order

// CustomizationChoice represents an order customization.
type CustomizationChoice struct {
	APIKey  string  `json:"apiKey"`
	Details string  `json:"details"`
	Price   float32 `json:"price"`
}

// Item represents an item for an order.
type Item struct {
	APIKey               string                `json:"apiKey"`
	Name                 string                `json:"name"`
	Comments             string                `json:"comments"`
	Price                float32               `json:"price"`
	TotalPrice           float32               `json:"totalPrice"`
	CustomizationChoices []CustomizationChoice `json:"customizationChoices"`
}

// Order represents a restaurant order.
type Order struct {
	APIKey           string `json:"apiKey"`
	ID               int    `json:"id"`
	DatePlaced       int64  `json:"datePlaced"`
	Method           string `json:"method"`
	Payment          string `json:"payment"`
	RestaurantAPIKey string `json:"restaurantAPIKey"`
	RecipientAPIKey  string `json:"recipientAPIKey"`
	Items            []Item `json:"items"`
}

// APIKey is a helper for the APIKey embedded in a new order request.
type APIKey struct {
	APIKey string
}

// NewOrder represents a json to place an order request to the API.
type NewOrder struct {
	RestaurantAPIKey string `json:"restaurantApiKey"`
	Items            []Item `json:"items"`
	Method           string `json:"method"`
	Payment          string `json:"payment"`
	Test             bool   `json:"test"`
	Card             APIKey `json:"card"`
	Address          APIKey `json:"address"`
	Recipient        APIKey `json:"recipient"`
}
