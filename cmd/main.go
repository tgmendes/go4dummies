package main

import "fmt"

func main() {
	fmt.Println(0.1 + 0.2)
	// db, err := bolt.Open("restaurant.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// tx, err := db.Begin(true)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer tx.Rollback()

	// _, err = tx.CreateBucketIfNotExists([]byte("Restaurant"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err = tx.Commit(); err != nil {
	// 	log.Fatal(err)
	// }

	// // db.Update(func(tx *bolt.Tx) error {
	// // 	b := tx.Bucket([]byte("Restaurant"))
	// // 	err := b.Put([]byte("answer"), []byte("42"))
	// // 	return err
	// // })

	// db.View(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("Restaurant"))
	// 	v := b.Get([]byte("s"))
	// 	fmt.Printf("The answer is: %s\n", v)
	// 	return nil
	// })
}
