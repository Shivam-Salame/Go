package main

import "fmt"

func main() {
	fmt.Println("Structs in go")
	// no inheritance in golang: no super or parent

	shivam := User{"shivam", "shivam.salame@google.com", true, 24}
	fmt.Println(shivam)
	fmt.Printf("shivam details are: %+v\n", shivam)
	fmt.Println("name: %v, email: %v, status: %v, age: %v", shivam.Name, shivam.Email, shivam.Status, shivam.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
