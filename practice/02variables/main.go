package main

import "fmt"

// starting with capital letter will be public variable
const LoginToken string = "public variable"

func main() {
	// string
	var username string = "shivam"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	// bool
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// uint
	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)

	// float
	var smallFloat float32 = 255.2333333333333338984392894832
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default value and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "shivam.com"
	fmt.Println(website)

	// no var style, walrus operator cant be used outside fuction,
	// if you have to use outside main function then var numberOfUser int = 300000
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
}
