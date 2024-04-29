package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3,5)
	fmt.Println("Result is:",result)

	proResult,proMessage := proAdder(2,5,8,1)
	fmt.Println("Pro result is:",proResult)
	fmt.Println("Pro message is:",proMessage)



}

func greeter()  {
	fmt.Println("Namaste from golang")
}

func adder(valOne int, valTwo int) int{
	return valOne + valTwo
}

//  Variadic functions
func proAdder(values ...int) (int,string) {
	total := 0
	for _,value := range values{
		total += value
	}
	return total, "Hi pro result function"
}
