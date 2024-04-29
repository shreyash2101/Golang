package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in golang")

	// no inheritance in golang; no super or parent

	shreyash := User{"Shreyash", "shreyash@go.dev", true, 24}
	fmt.Println(shreyash)
	fmt.Printf("shreyash details are: %v\n", shreyash)
	fmt.Printf("shreyash details are: %+v\n", shreyash)
	fmt.Printf("Name is %v and email is %v.", shreyash.Name,shreyash.Email)
}

type User struct {
	Name 	string
	Email 	string
	Status 	bool
	Age 	int
}