package main

import (
	"fmt"
)

func main() {
	var firstName string = "Usman"
	var lastName string = "Zahid"
	var fullName string = firstName + " " + lastName
	
	fmt.Println("Hello there", fullName)
}

func addNumbers(x int, y int) int {
	return x + y
}