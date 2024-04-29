package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in golang")

	var fruitList = []string{"Apple", "Banana", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Guava")
	fmt.Println(fruitList)

	fruitList=append(fruitList[1:3])
	fmt.Println(fruitList)

	fruitList=append(fruitList[:4])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	// not allowed
	// highScores[4] = 777

	// allowed
	highScores = append(highScores, 555,666,321)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from slice based on index
	var courses = []string{"ReactJS", "JavaScript", "Swift", "Python", "Ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}