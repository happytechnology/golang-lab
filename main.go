package main

import "fmt"

func main() {
	var username string = "name"
	var password int = 12345

	fmt.Println("Auth: Basic", username+":"+password)
}
