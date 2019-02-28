package restaurant

type DatabaseConnector interface {
	Add(entity []byte) ([]byte, error)
	List() ([]byte, error)
	View(id int) ([]byte, error)
	Update(id int, entity []byte) error
	Delete(id int) error
}

// Fetcher represents an interface to fetch web pages.
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

