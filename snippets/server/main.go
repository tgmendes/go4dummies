package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("New connection from ", r.RemoteAddr)
		w.Write([]byte("Hello, world!"))
	})

	log.Println("listening to connections on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
