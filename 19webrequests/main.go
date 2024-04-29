package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://go.dev/"

func main() {

	fmt.Println("Welcome to web request in golang")

	resp, err := http.Get(url)

	if err!=nil{
		panic(err)
	}

	fmt.Printf("Response is of type: %T", resp)

	defer resp.Body.Close() // it is caller's responsibility to close the connection

	databytes,err := io.ReadAll(resp.Body)

	if err!=nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)

}