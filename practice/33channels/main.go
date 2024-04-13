package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels")

	// myChannel := make(chan int)

	// will trow error : all goroutines are asleep - deadlock!
	// myChannel <- 5
	// fmt.Println(<-myChannel)

	// myCh := make(chan int, 1) // if second argument we add 1 so it wont throw deadlock even if 2nd listener is not present, its called buffer channel we are having many buffer
	myCh := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup){
		fmt.Println(<-myCh)
		fmt.Println(<-myCh) // adding the second listener so deadlock wont arise
		wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup){
		myCh <- 5
		myCh <-6 // if we pass 2 values and there is not listener for the second value it will trow deadlock error
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}

