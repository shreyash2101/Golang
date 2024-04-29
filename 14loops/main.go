package main

import "fmt"

func main() {

	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d:=0; d<len(days); d++{
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days{
		fmt.Printf("Index is %v and value is %v\n", index,day)
	}

	rougueValue := 1

	for rougueValue<10 {

		// Simply jumps to test
		if rougueValue ==2 {
			goto test
		}

		// loop breaks as soon as 5 matches
		// if rougueValue ==5 {
		// 	break
		// }

		// loop skips when 5 matches
		// if rougueValue ==5 {
		// 	rougueValue++
		// 	continue
		// }

		fmt.Println("Value is:", rougueValue)
		rougueValue++
	}

	test:
		fmt.Println("Jumping at test")

}