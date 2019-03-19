package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	c := make(chan string)

	go greeter("Luke", c, 1000)   // HL
	go greeter("Leia", c, 10)     // HL
	go greeter("Obi-Wan", c, 400) // HL

	for n := 0; n < 3; n++ {
		// Wait until a channel is ready.
		select { // HL
		case res := <-c:
			fmt.Println(res)
		}
	}
}

func greeter(person string, c chan string, delay int) {
	time.Sleep(time.Duration(delay) * time.Millisecond)
	c <- fmt.Sprintf("Hello, %s", person) // HL
}

// END OMIT
