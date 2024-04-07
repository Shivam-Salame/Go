package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("make post request")

	PerformPostRequest()
}

func PerformPostRequest() {
	const myUrl = "http://localhost:3000/"

	// fake json payload

	requestBody := strings.NewReader(`
		{
			"courseName":"Lets go with golang",
			"price":"10",
			"platform":"Go lang dev"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
