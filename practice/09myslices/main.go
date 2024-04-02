package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("slices:")

	var fruitList = []string{"apple", "tomato", "beans"}
	fmt.Printf("Type of fruitList %T\n", fruitList)
	fmt.Println("fruitList:", fruitList)

	fruitList = append(fruitList, "banana", "kiwi")
	fmt.Println("fruitList:", fruitList)

	fruitList = append(fruitList[1:]) // start: end
	fmt.Println("fruitList:", fruitList)

	fruitList = append(fruitList[1:3]) // start: end+1
	fmt.Println("fruitList:", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 1
	highScores[1] = 9
	highScores[2] = 3
	highScores[3] = 4
	// highScores[4] = 5 // this will give error here outof bound

	fmt.Println("highScores:", highScores)

	highScores = append(highScores, 6, 7, 8)
	fmt.Println("highScores:", highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println("highScores:", highScores)

	// remove a value from slices based on index

	var courses = []string{"c++", "python", "go", "javascript", "c"}
	fmt.Println("courses:", courses)

	var index int = 2
	// we can also use append to remove the values
	courses = append(courses[:index], courses[index + 1:]... )
	fmt.Println("courses:", courses)
}
