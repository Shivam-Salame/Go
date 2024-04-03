package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://dummyjson.com/products/1"

func main() {
	fmt.Println("handling web requests in go")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)

}
