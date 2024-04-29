package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition in golang")
	
	wg := &sync.WaitGroup{}

	var score = []int{0}

// ---------------------------------------------
// Race Condition
// Use go run --race . to check if it has race condition or not

	// wg.Add(3) // as we have three go routines below so we are adding 3 to the waitgroup
	// go func(wg *sync.WaitGroup) {
	// 	fmt.Println("One R")
	// 	score = append(score, 1)
	// 	wg.Done()
	// }(wg)

	// go func(wg *sync.WaitGroup) {
	// 	fmt.Println("Two R")
	// 	score = append(score, 2)
	// 	wg.Done()
	// }(wg)

	// go func(wg *sync.WaitGroup) {
	// 	fmt.Println("Three R")
	// 	score = append(score, 3)
	// 	wg.Done()
	// }(wg)

	// wg.Wait()
	// fmt.Println(score)

// ------------------------------------------------------------
// Using simple mutex to overcome race condition

	// mut := &sync.Mutex{}

	// wg.Add(3)
	// go func(wg *sync.WaitGroup, m *sync.Mutex) {
	// 	fmt.Println("One R")
	// 	mut.Lock()
	// 	score = append(score, 1)
	// 	mut.Unlock()
	// 	wg.Done()
	// }(wg, mut)

	// go func(wg *sync.WaitGroup, m *sync.Mutex) {
	// 	fmt.Println("Two R")
	// 	mut.Lock()
	// 	score = append(score, 2)
	// 	mut.Unlock()
	// 	wg.Done()
	// }(wg, mut)

	// go func(wg *sync.WaitGroup, m *sync.Mutex) {
	// 	fmt.Println("Three R")
	// 	mut.Lock()
	// 	score = append(score, 3)
	// 	mut.Unlock()
	// 	wg.Done()
	// }(wg, mut)

	// wg.Wait()
	// fmt.Println(score)

// ------------------------------------------------------------
// Using Read Write mutex to overcome race condition

	mut := &sync.RWMutex{}

	wg.Add(4)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Four R")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)

}
