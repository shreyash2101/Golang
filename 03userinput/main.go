package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin) 
	fmt.Println("Enter your input")

	// comma ok syntax || error ok syntax

	input, _ := reader.ReadString('\n')
	// input, err := reader.ReadString('\n')
	fmt.Println("Thanks for the value", input)
	fmt.Printf("Type of the input is: %T", input)
}