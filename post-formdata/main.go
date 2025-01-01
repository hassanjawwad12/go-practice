package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const URL string = "https://reqres.in/api/users"

func main() {

	// Create the formdata
	data := url.Values{}
	data.Add("name", "Hassan Jawwad")
	data.Add("job", "ASE")

	// Post the formdata
	resp, err := http.PostForm(URL, data)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.StatusCode)

	responseData := make([]byte, 4096) // Allocate a buffer for the response
	n, _ := resp.Body.Read(responseData)
	content := string(responseData[:n])
	fmt.Println("Response Body:", content)
}
