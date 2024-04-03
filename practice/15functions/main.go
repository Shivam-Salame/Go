package main

import "fmt"

func main() {
	fmt.Println("function in go")
	namastey()

	result := adder(1, 3)
	fmt.Println("sum of two numbers:", result)

	proResultInt, proResultString := proAdder(1, 2, 4, 5, 6, 7)
	fmt.Printf("proAdder total is: %v\nproResultString is: %v", proResultInt, proResultString)
}

func namastey() {
	fmt.Println("namastey from golang")
}

// explicitly we have to mention return type
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// multiple values
func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "hello from proAdder"
}

// anonymous function
// func () {
// 	fmt.Println("anonymous function")
// } ()

// iife (immediately invoked function expression)
