package main

import "fmt"

func main() {
	fmt.Println("pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is:", ptr) //null value

	myNumber := 25
	var ptr = &myNumber
	fmt.Println("Value of pointer is:", *ptr)

	*ptr = *ptr + 2
	fmt.Println("new value is:", myNumber)

}
