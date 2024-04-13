package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}

	wg.Add(3)
	// go func (w *sync.WaitGroup) {
	// 	fmt.Println("One R")
	// 	score = append(score, 1)
	// 	wg.Done()
	// }(wg)

	// //wg.Add(1)
	// go func (w *sync.WaitGroup) {
	// 	fmt.Println("Two R")
	// 	score = append(score, 2)
	// 	wg.Done()
	// }(wg)

	// go func (w *sync.WaitGroup) {
	// 	fmt.Println("Three R")
	// 	score = append(score, 3)
	// 	wg.Done()
	// }(wg)

	// cleared race codition by adding lock to process : go run --race .
	// in case of read lock unlock use it sam eplace where we have used not on resource like var score = []int{0}
	go func(w *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	//wg.Add(1)
	go func(w *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(w *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()

	// it doesn't depend
	fmt.Println(score)
}
