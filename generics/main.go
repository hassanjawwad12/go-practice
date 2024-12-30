package main

import "fmt"

// it infers the type but it is not a good practice to use any/interface
// func printSlice[T interface{}](items []T)
func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type stack[T int | string | bool] struct {
	elements []T
}

func main() {
	nums := []int{1, 2, 3}
	names := []string{"Hassan", "Ali"}
	printSlice(names)
	printSlice(nums)

	myStack := stack[int]{
		elements: []int{4, 5, 6},
	}
	myStack2 := stack[string]{
		elements: []string{"hassan", "nauman", "ali"},
	}
	fmt.Println(myStack)
	fmt.Println(myStack2)
}
