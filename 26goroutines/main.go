package main

import (
	"fmt"
	"time"
)

func main() {

	// greeter("Hello")
	// greeter("world")

	// go greeter("Hello")
	// greeter("world")

// -----------------------------------

	// greeter2("Hello")
	// greeter2("world")

	go greeter2("Hello")
	greeter2("world")

}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func greeter2(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond) // used to make sure the go routine is executed before exiting the main method
		fmt.Println(s)
	}
}
