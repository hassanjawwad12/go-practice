package main

import "fmt"

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
func main() {
	num := 1
	changeNum(num)

	fmt.Println("value does not change ", num)

	fmt.Println("Mem Address ", &num)

	changeNumber(&num)
	fmt.Println("value does change ", num)

}
