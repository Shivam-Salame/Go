package main

import (
	"encoding/json"
	"fmt"
)

// not public
type course struct {
	Name     string
	Price    int
	Platform string
	Password string
	Tags     []string
}

type indentCourses struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("creating json data")

	EncodeJson()
}

func EncodeJson() {
	courses := []course{
		{"React Bootcamp", 29, "shivam.com", "heyo", []string{"hii", "hello", "ladies washroom"}},
		{"angular Bootcamp", 10, "salame.com", "heyo", []string{"hii", "hello", "ladies washroom"}},
		{"Vue Bootcamp", 30, "bebo.com", "heyo", []string{"hii", "hello", "ladies washroom"}},
	}

	indentCourses := []course{
		{"React Bootcamp", 29, "shivam.com", "heyo", []string{"hii", "hello", "ladies washroom"}},
		{"angular Bootcamp", 10, "salame.com", "heyo", []string{"hii", "hello", "ladies washroom"}},
		{"Vue Bootcamp", 30, "bebo.com", "heyo", nil},
	}

	// package this data as json data
	finalJson, err := json.Marshal(courses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

	// indent json format ( jsonData, "prefix-string", "space-aarangement")
	indentFinalJson, err := json.MarshalIndent(indentCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", indentFinalJson)
}
