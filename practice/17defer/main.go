package main

import "fmt"

func main() {
	// in the reverse order they are defered
	defer fmt.Println("defer one")
	defer fmt.Println("defer two")
	defer fmt.Println("defer three")
	fmt.Println("hello world")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("defer func value: ", i)
	}
}
