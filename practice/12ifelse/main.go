package main

import "fmt"

func main() {
	fmt.Println("if else in golang")

	loginCount := 25
	var result string
	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 && loginCount < 20 {
		result = "premium user"
	} else {
		result = "bussiness user"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("even no.")
	} else {
		fmt.Println("odd no.")
	}

	if num := 3; num < 10 {
		fmt.Println("no. smaller than 10")
	} else {
		fmt.Println("no. greater than 10")
	}

	// if err != nil {

	// }
}
