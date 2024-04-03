package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch case in go")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value 1")
	case 2:
		fmt.Println("dice value 2")
	case 3:
		fmt.Println("dice value 3")
		fallthrough
	case 4:
		fmt.Println("dice value 4")
		fallthrough
	case 5:
		fmt.Println("dice value 5")
	case 6:
		fmt.Println("dice value 6")
	default:
		fmt.Println("what was that")

	}
}
