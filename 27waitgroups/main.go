package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //pointer

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
	// 	 getStatusCode(web)
	// }

// -----------------------------------------

// This is faster in real time as it runs concurrently in goroutines
// WaitGroup is used to make sure we wait for go routine to get executed before exiting the main method

	for _, web := range websitelist {
		go getStatusCodeWithWaitGroup(web)
		wg.Add(1)
	}

	wg.Wait()

}

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	}else{
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}

func getStatusCodeWithWaitGroup(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
