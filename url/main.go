package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://jsonplaceholder.typicode.com/todos/1"

func main() {

	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()

	fmt.Printf("the type of query params are: %T\n", qparams)

	for _, val := range qparams {
		fmt.Println("The params are: ", val)
	}
}
