package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels is the way how goroutine talk to each other")

	// myChannel := make(chan int)

	// will trow error : all goroutines are asleep - deadlock!
	// myChannel <- 5
	// fmt.Println(<-myChannel)

	// myCh := make(chan int, 1) // if second argument we add 1 so it wont throw deadlock even if 2nd listener is not present, its called buffer channel we are having many buffer
	myCh := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	// to make read only goroutine "<-chan"
	go func(ch chan int, wg *sync.WaitGroup){
		// close(myCh) will give error here it is readonly goroutine

		fmt.Println(<-myCh)
		fmt.Println(<-myCh) // adding the second listener so deadlock wont arise

		val, isChannelOpen := <-myCh
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(myCh, wg)
	
	// to make send only goroutine "chan<-"
	go func(ch chan int, wg *sync.WaitGroup){
		// close(myCh) // if we close the channel before passing values it will throw an error: send on closed channel,
		// if we don't pass any values it takes by default 0, if we pass 0 value before closing connection it will work
		myCh <- 5
		myCh <- 6 // if we pass 2 values and there is not listener for the second value it will trow deadlock error
		myCh <- 0
		close(myCh)
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}

