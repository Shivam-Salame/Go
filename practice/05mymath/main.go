package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	"crypto/rand"
	// "time"
)

func main() {
	fmt.Println("crypto, math, random number")

	var myNumOne int = 2
	var myNumTwo float64 = 4.56

	fmt.Println("The sum is:", myNumOne + int(myNumTwo))

	// random 
	// exclde the digit 5
	// fmt.Println(rand.Intn(5))

	// it will provide truly random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1) // so it will give values 0-5

	// random from crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5)) // 0-5
	fmt.Println(myRandomNum)

}
