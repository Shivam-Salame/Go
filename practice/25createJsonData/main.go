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

type indentCourse struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"` // will never be encoded
	Tags     []string `json:"tags,omitempty"` // field omitted when it is set to its zero-value
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

	indentCourses := []indentCourse{
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
