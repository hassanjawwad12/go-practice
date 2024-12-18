package main

import "fmt"

//only for loop in go , no while loop
func main() {
	//while loop is run using for loop keyword
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	//infinite loop if you remove the break
	for {
		fmt.Println("loop")
		break
	}

	//classic for loop
	for j := 7; j <= 9; j++ {
		//this will not print 8
		if j == 8 {
			continue
		}
		fmt.Println(j)
	}

	//range loop
	for k := range 5 {
		fmt.Println(k)
	}
}
