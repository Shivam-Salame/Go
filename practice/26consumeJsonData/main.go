package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string
	Price    int
	Platform string
	Password string
	Tags     []string
}

func main() {
	fmt.Println("Decoding json")
	DecodingJson()
}

func DecodingJson() {

	// adding , at the end of tags line will anser JSON was not valid otherwise removal of , will result valid JSON and data
	jsonDataFromWeb := []byte(`
	{
		"coursename": "Reactjs",
		"Price": 299,
		"website": "lco.in",
		"tags": ["web-dev","js"] 
	}
	`)
	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse) // we don't want to pass data as copy
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		// order is not maintained in maps
		fmt.Printf("key is %v, value is %T, type is %v\n", k, v, v)
	}
}
