package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

// ------------------------------------------------------------
// Deadlock situation - we cannot simply create a channel and use it without goroutine
// Channel can only listen if someone is throwing a value to it
// Channel can only throw a value only if someone is listening to it

	// myCh := make(chan int)

	// myCh <- 5
	// fmt.Println(<-myCh)

// ------------------------------------------------------------
// Removing Deadlock
// By using Gorutines we overcome the deadlock as both listening and throwing of values happens concurrently

	// myCh := make(chan int)
	// wg := &sync.WaitGroup{}

	// wg.Add(2)
	// go func(ch chan int, wg *sync.WaitGroup) {

	// 	fmt.Println(<-myCh)
	// 	wg.Done()
	// }(myCh, wg)

	// go func(ch chan int, wg *sync.WaitGroup) {
	// 	myCh <- 6
	// 	wg.Done()
	// }(myCh, wg)

	// wg.Wait()

// ------------------------------------------------------------
// Buffered Channel

	// myCh := make(chan int, 2) // Buffered channel
	// wg := &sync.WaitGroup{}

	// wg.Add(2)
	// go func(ch chan int, wg *sync.WaitGroup) {
	// 	for i:=0;i<2;i++ {
	// 		fmt.Println(<-myCh)
	// 	}
	// 	wg.Done()
	// }(myCh, wg)

	// go func(ch chan int, wg *sync.WaitGroup) {
	// 	myCh <- 6
	// 	myCh <- 7
	// 	wg.Done()
	// }(myCh, wg)

	// wg.Wait()

// ------------------------------------------------------------
// Closed Channel & Uni directional channels

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// R ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChanelOpen := <-myCh

		fmt.Println(isChanelOpen)
		fmt.Println(val)

		//fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)

	// send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		// myCh <- 6
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
