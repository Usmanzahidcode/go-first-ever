package main

import (
	"fmt"
)

func main() {	
	name := readName()
	
    fmt.Println("Hello,", name)
}

func readName() string {
	var name string
	
	fmt.Print("Enter your name: ")
	_, err := fmt.Scan(&name)
	
	if err != nil {
        fmt.Println("Error reading name:", err)
    }
	
	return name
}
