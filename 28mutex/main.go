package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup //pointer
var mut sync.Mutex    // pointer

func main() {

	websitelist := []string{
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://test.dev",
		"https://github.com",
	}

// -----------------------------------------

	// for _, web := range websitelist {
	// 	go getStatusCode(web)
	// 	wg.Add(1)
	// }

	// wg.Wait()
	// fmt.Println(signals)

// -----------------------------------------
// No significant difference would be seen in this case as we have only one goroutine but we have better example in next module

	for _, web := range websitelist {
		go getStatusCodeWithMutex(web)
		wg.Add(1)
	}


	wg.Wait()
	fmt.Println(signals)
}


func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		signals = append(signals, endpoint)

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}

func getStatusCodeWithMutex(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		
		// if multiple goroutines are accessing this global variable and doing read write operation then this mutex will make sure that the memory is locked for this particular goroutine and no other goroutines access it until it is unlocked

		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
