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

func fibonacci() func() int {
	a, b := 0, 1

	return func() int {
		current := a
		a, b = b, a+b
		return current
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

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
