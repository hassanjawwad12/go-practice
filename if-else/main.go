package main

import "fmt"

func main() {

	age := 18
	if age >= 18 {
		fmt.Println("You are eligible to vote")
	} else if age >= 16 {
		fmt.Println("You are eligible to vote but with parent's permission")
	} else {
		fmt.Println("You are not eligible to vote")
	}

	var role= "admin"
	var permission =false
	if role == "admin" || permission {
		fmt.Println("You have access to all the resources")
	} else {
		fmt.Println("You have limited access")
	}

	//we can declare variable inside if condition
	if number := 10; number%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}
	

	//no ternary operator in go, use if else
}
