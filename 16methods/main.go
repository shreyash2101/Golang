package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in golang")

	// no inheritance in golang; no super or parent

	shreyash := User{"Shreyash", "shreyash@go.dev", true, 24}
	fmt.Println(shreyash)
	fmt.Printf("shreyash details are: %v\n", shreyash)
	fmt.Printf("shreyash details are: %+v\n", shreyash)
	fmt.Printf("Name is %v and email is %v.\n", shreyash.Name,shreyash.Email)

	// Methods
	shreyash.GetStatus()
	shreyash.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", shreyash.Name,shreyash.Email) // a different copy of the original value is passed to the method

}

type User struct {
	Name 	string
	Email 	string
	Status 	bool
	Age 	int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}