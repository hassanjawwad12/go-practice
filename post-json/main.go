package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const URL string = "https://reqres.in/api/users"

type RequestBody struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

func main() {

	requestBody := RequestBody{
		Name: "Hassan Jawwad",
		Job:  "ASE",
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	request, error := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
	if error != nil {
		panic(error)
	}
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

	var result map[string]any
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}
	fmt.Println("Response:", result)
}