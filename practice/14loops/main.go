package main

import "fmt"

func main() {
	fmt.Println("loops in golang")

	days := []string{"sunday", "monday", "tuesday", "wednesday", "thursday"}
	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	fmt.Println("-----------------")

	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("-----------------")

	// here we can use _ in place of index
	for index, day := range days {
		fmt.Printf("index is %v day is %v\n", index, day)
	}

	fmt.Println("-----------------")

	rougeValue := 1

	for rougeValue < 10 {
		if rougeValue == 7{
			goto lco
		}

		if rougeValue == 3 {
			rougeValue++
			continue
		}

		fmt.Println("rougeValue is:", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("value is run using goto")
}
