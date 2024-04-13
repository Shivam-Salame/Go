package main

import (
	"fmt"
	"net/http"
	"sync"
)

// waitgroup 
// wait until goroutine finishes its work
var wg sync.WaitGroup // pointer
var mut sync.Mutex //pointer

var signals = []string{"test"}

func main() {
	// the "go" keywords is used for goroutine
	// go greeter("hello")
	// greeter("shivam")

	websiteList := []string{
		"https://loc.dev",
		"https://go.dev",
		"https://google.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		// add this to enable waitgroup, we entered 1 because only one call is happening
		wg.Add(1)
	}

	// this doesn't allow main method to be finished because my friends waitgroups are still running
	wg.Wait()

	fmt.Println(signals)
}

// func greeter(s string) {
// for i := 0; i < 3; i++ {
// time.Sleep(5 * time.Millisecond)
// fmt.Println(s)
// }
// }

func getStatusCode(endpoint string) {
	// pass the signal at the end that it is finished
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPs in endpoint")
	} else {
		// operation while reading and writing
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf(" %d status code for %s\n", res.StatusCode, endpoint)

	}
}

// mutex will lock the memory until goroutine finished its work 
// it has one more brother read write mutex, in that it allows any one to read memory but while writing 