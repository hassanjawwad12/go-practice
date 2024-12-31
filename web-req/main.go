package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const url = "https://jsonplaceholder.typicode.com/todos"

func main() {

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("response is of type: %T\n", response)

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(data)

	newFile, err := os.Create("example.txt")
	if err != nil {
		panic(err)
	}

	defer newFile.Close()

	_, err = io.WriteString(newFile, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("Content written to file successfully!")

}
