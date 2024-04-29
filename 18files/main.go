package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a file"

	file,err := os.Create("./myfile.txt")

	// if err!=nil {
	// 	panic(err)
	// }
	checkNilError(err)

	length, err := io.WriteString(file,content)

	checkNilError(err)

	fmt.Println("Length is:", length)

	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(fileName string)  {
	dataByte,err := os.ReadFile(fileName)

	checkNilError(err)

	fmt.Println("Text data inside the file is \n", string(dataByte))

}

func checkNilError(err error)  {
	if err!=nil {
		panic(err)
	}
}