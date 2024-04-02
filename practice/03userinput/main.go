package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user imput"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("rating:")

	// comma ok || error ok ( input, error ) or ( input, _ ) if we dont want error or ( _, error ) if we only want error
	input, _ := reader.ReadString('\n')
	fmt.Println("rated:", input) // this automatically add newline at end
	fmt.Printf("Type fo this rating string is %T \n", input)
}
