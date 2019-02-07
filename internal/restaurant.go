package restaurant

type DatabaseConnector interface {
	Add(entity []byte) ([]byte, error)
	List() ([]byte, error)
	View(id int) ([]byte, error)
	Update(id int, entity []byte) error
	Delete(id int) error
}
