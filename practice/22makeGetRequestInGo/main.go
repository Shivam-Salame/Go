package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("make get request")

	PerformGetRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:3000/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	// defer response.Body.Close()

	fmt.Println("Status code:", response.Status)
	fmt.Println("content length is:", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	
	// first way using library string.Builder
	fmt.Println("Byte count is: ", byteCount)
	
	
	// second way
	fmt.Println(content)
	fmt.Println(string(content))

}
