package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList[0] = "Apple"
	fruitList[1]="Banana"
	fruitList[3]= "Peach"

	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("Fruit list is:", len(fruitList))

	var vegList = [3]string{"Potato", "beans", "mushroom"}

	fmt.Println("Veggies list is:", vegList)
	fmt.Println("Veggies list is:", len(vegList))


}