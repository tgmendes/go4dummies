package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i

	fmt.Println("memory address of p: ", p)
	fmt.Println("value of p: ", *p)
	*p = 21
	fmt.Println("new value of i:", i)

	p = &j
	*p = *p / 37
	fmt.Println("value of j:", j)
}
