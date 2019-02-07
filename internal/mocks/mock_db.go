package mocks

type DB struct {
	AddFn          func(entity []byte) ([]byte, error)
	AddFnCalled    int
	ListFn         func() ([]byte, error)
	ListFnCalled   int
	ViewFn         func(id int) ([]byte, error)
	ViewFnCalled   int
	UpdateFn       func(id int, entity []byte) error
	UpdateFnCalled int
	DeleteFn       func(id int) error
	DeleteFnCalled int
}

func (db *DB) Add(entity []byte) ([]byte, error) {

	db.AddFnCalled++
	return db.AddFn(entity)
}
func (db *DB) List() ([]byte, error) {
	db.ListFnCalled++
	return db.ListFn()
}
func (db *DB) View(id int) ([]byte, error) {
	db.ViewFnCalled++
	return db.ViewFn(id)
}
func (db *DB) Update(id int, entity []byte) error {
	db.UpdateFnCalled++
	return db.UpdateFn(id, entity)
}
func (db *DB) Delete(id int) error {
	db.DeleteFnCalled++
	return db.DeleteFn(id)
}
