package main

import "fmt"

func main() {
	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[3] = "kiwi"

	fmt.Println("fruitList:", fruitList) // extra space in third elements as it was not present in array
	fmt.Println("len of fruitList array:", len(fruitList))

	var vegList = [3]string{"potato", "okra", "beans"}
	fmt.Println("vegList:",vegList)
}
