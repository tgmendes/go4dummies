package customer

// Customer represents a model of a restaurant customer.
type Customer struct {
	ID    int
	Name  string
	Email string
}

type NewCustomer struct {
	Name  string
	Email string
}
