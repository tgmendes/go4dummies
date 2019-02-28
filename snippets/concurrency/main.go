package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	c := make(chan string)

	go greeter("Luke", c, 1000)
	go greeter("Leia", c, 10)
	go greeter("Obi-Wan", c, 400)

	// go func() {
	for n := 0; n < 3; n++ {
		select {
		case res := <-c:
			fmt.Println(res)
		}
	}
}

func greeter(person string, c chan string, delay int) {
	time.Sleep(time.Duration(delay) * time.Millisecond)
	c <- fmt.Sprintf("Hello, %s", person)
}

// END OMIT
