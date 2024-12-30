package main

import "fmt"

//array is a fixed length sequence of elements of a specific type
func main() {

	var nums [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(len(nums))

	var numbers [4]int

	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	//numbers[3]=4
	//it will print 0 for the indexes which are not assigned any integer value
	fmt.Println(numbers)

	var vals [4]bool
	vals[2] = true
	fmt.Println(vals)

	//simpler syntax
	prices := [3]float64{1.1, 2.2, 3.3}
	fmt.Println(prices[2])

	//2D arrays
	two := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(two)
}
