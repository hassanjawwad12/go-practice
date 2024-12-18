package main

import "fmt"

// A function that returns a closure
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	counter := makeCounter()

	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	anotherCounter := makeCounter()
	fmt.Println(anotherCounter())
	fmt.Println(anotherCounter())
}
