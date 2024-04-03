package main

import (
	"fmt"
	"net/url"
)

const uri = "https://dummyjson.com/products/1"

func main() {
	fmt.Println("handling url in go")
	fmt.Println(uri)

	// parsing
	result, _ := url.Parse(uri)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)
	// fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
