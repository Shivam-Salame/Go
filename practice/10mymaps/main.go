package main

import "fmt"

func main() {
	fmt.Println("Maps in go")

	// map[key]value
	languages := make(map[string]string)

	// by default sorted
	languages["JS"] = "javascript"
	languages["NJ"] = "nodejs"
	languages["C"] = "c"
	languages["GO"] = "golang"
	languages["PY"] = "python"

	fmt.Println("List of all languages:", languages)
	fmt.Println("GO shorts for:", languages["GO"])

	delete(languages, "NJ")
	fmt.Println("List of all languages:", languages)

	// loops
	for key, value := range languages {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}
	for _, value := range languages {
		fmt.Printf("value: %v\n", value)
	}
}
