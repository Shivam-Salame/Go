package main

import "fmt"

func main() {
	fmt.Println("Structs in go")
	// no inheritance in golang: no super or parent

	shivam := User{"shivam", "shivam.salame@google.com", true, 24}
	fmt.Println(shivam)
	fmt.Printf("shivam details are: %+v\n", shivam)
	fmt.Printf("name: %v, email: %v, status: %v, age: %v\n", shivam.Name, shivam.Email, shivam.Status, shivam.Age)

	shivam.GetStatus()
	shivam.GetEmail()
	fmt.Printf("shivam details are: %+v\n", shivam)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	//oneAge int  // this property wont be exportable because first letter is not capital
}

func (u User) GetStatus(){
	fmt.Println("status is:", u.Status)
}

// this creates a copy for every new object
func (u User) GetEmail(){
	u.Email = "shivam.salame@ril.com"
	fmt.Println("Email is:", u.Email)
}