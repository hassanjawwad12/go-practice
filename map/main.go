package main

import (
	"fmt"
	"maps"
)

//map are basically hash ,obj,dict
func main () {

	//key type and value type
	m:= make(map[string]string)

	//setting an element 
	m["name"] = "John"
	m ["age"] = "25"

	//getting an element
	fmt.Println("Printing first map")
	fmt.Println(m["name"],m["age"])
	fmt.Println(m)

	//if key value does not exist, it will return empty string
	//fmt.Println (m["dob"])

	//delete an element
	delete(m,"age")

	//check if key exists  _ have the value of the key
	_, ok := m["name"]
	if ok {
		fmt.Println("Key exists")
	} else {
		fmt.Println("Key does not exist")
	}
	
	k:= make(map[string]int)
	k["age"] = 25
	fmt.Println("Printing second map")
	fmt.Println(k["age"])
	fmt.Println(len(k))
	clear(k)

	l:= map[string]int{"age":25,"dob":1995}
	fmt.Println("Printing third map")
	fmt.Println(l["age"])

	//check if 2 maps are equal
	fmt.Println("Checking if 2 maps are equal")
	fmt.Println(maps.Equal(k,l))
}

