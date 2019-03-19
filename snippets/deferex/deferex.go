package main

import "fmt"

func main() {
	defer doSomething()

	fmt.Println("I will be called first")
}

func doSomething() {
	fmt.Println("I will be called at the end")
}
