package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const URL string = "https://reqres.in/api/users"

func main() {
	var jsonData = []byte(`{
		"name": "Hassan Jawwad",
		"job": "ASE"
	}`)

	request, error := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	fmt.Println("Response Status:", response.Status)
	fmt.Printf("\n")

	fmt.Println("response Headers:", response.Header)
	fmt.Printf("\n")

	body, _ := io.ReadAll(response.Body)
	content := string(body)

	fmt.Println("Response Body:", content)

}
