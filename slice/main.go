package main

import (
	"fmt"
	"slices"
)

// slice is a dynamically sized flexible view of elements of an array
// most used construct in go
func main() {

	//uninitialized slice is nil
	var new []int
	fmt.Println(new == nil)

	//capacity growing with each element
	random := []int{}
	random = append(random, 1)
	random = append(random, 2)
	fmt.Println(random)
	fmt.Println(cap(random))

	//this 2 is the initial size
	//5 is capacity
	var nums = make([]int, 0, 5)
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)
	nums = append(nums, 6)

	//capacity of the slice ( max number of elements)
	//capacity is doubled automatically
	fmt.Println("This is the nums 1:", nums)
	fmt.Println("This is the capacity of nums1:", cap(nums))

	//copy a slice
	var nums2 = make([]int, len(nums))
	copy(nums2, nums)
	fmt.Println("This is the nums 2", nums2)

	//slice operator
	var elements = []int{1, 2, 3, 4, 5}
	//excluding the 3
	fmt.Println(elements[1:3])
	fmt.Println(elements[:3])
	fmt.Println(elements[2:])

	//slice package
	var r1 = []int{1, 2, 3, 4, 5}
	var r2 = []int{6, 7, 8, 9, 10}
	fmt.Println(slices.Equal(r1, r2))

	//2d slices
	var two = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(two)

	// Create a 2D slice with 3 rows and 4 columns
	rows := 3
	cols := 4
	matrix := make([][]int, rows)

	// Initialize each row with a slice of length cols
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	// Manually assign values to each index
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[0][3] = 4

	matrix[1][0] = 5
	matrix[1][1] = 6
	matrix[1][2] = 7
	matrix[1][3] = 8

	matrix[2][0] = 9
	matrix[2][1] = 10
	matrix[2][2] = 11
	matrix[2][3] = 12

	// Print the 2D slice
	fmt.Println("2D slice with manually assigned values:")
	for i := 0; i < rows; i++ {
		fmt.Println(matrix[i])
	}

}
