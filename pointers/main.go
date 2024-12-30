package main

import (
	"fmt"
)

// passed by value
func changeNum(num int) {
	num = 5
	fmt.Println(num)
}

func changeNumber(num *int) {
	//dereferencing
	*num = 5
	fmt.Println(*num)
}

func change(num *[]int) []int {
	// Create a new slice and append the values from *num
	newSlice := append(*num, 5)
	return newSlice
}

func main() {
	num := 1
	changeNum(num)

	fmt.Println("value does not change ", num)

	fmt.Println("Mem Address ", &num)

	changeNumber(&num)
	fmt.Println("value does change ", num)

	numbers := []int{1, 2, 3, 4}
	result := change(&numbers)
	fmt.Println("In main function:", result)
	fmt.Println("Original slice remains unchanged:", numbers)
}
