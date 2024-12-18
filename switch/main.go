package main

import (
	"fmt"  
	"time"
)

func main () {
    //simple switch 
	i:=3

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Default")
	}


	//multiple condition switch 
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//type switch 
	whoAmI:= func (i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
	    case string:
			fmt.Println("I'm a string")		
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whoAmI("Hassan")
	whoAmI(1)
}

