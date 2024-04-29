package main

import "fmt"

const LoginToken string = "dhidqjknwj" // Public

// := is not allowed outside the method, so global variables cant be define using :=
// token := 134

// Allowed
var token1 = 45
var token2 int = 34

func main() {

	// String
	var username string = "Shreyash"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	// Boolean
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// Integer
	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// Float
	var smallFloat float64 = 256.3322224242424224
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "shreyash.in"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	// no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n", numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}