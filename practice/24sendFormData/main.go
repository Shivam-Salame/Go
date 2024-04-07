package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("make post request")
	PostFormRequest()
}

func PostFormRequest() {
	const myUrl = "http://localhost:3000/postform"

	// form-data
	data := url.Values{}
	data.Add("firstname", "shivam")
	data.Add("lastname", "salame")
	data.Add("email", "shivam@test.com")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
