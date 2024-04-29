package main

import (
	"fmt"
	"time"

	cryptoRandom "crypto/rand"
	"math/big"
	mathRandom "math/rand"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	var mynumberOne int = 2
	var mynumberTwo float64 = 4.5

	fmt.Println("The sum is: ", mynumberOne+int(mynumberTwo))

	//random from math

	mathRandom.Seed(time.Now().UnixNano()) // seed is mandatory for math/rand to work
	fmt.Println(mathRandom.Intn(5) + 1)

	//random from crypto

	myRandomNum, _ := cryptoRandom.Int(cryptoRandom.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)

}
