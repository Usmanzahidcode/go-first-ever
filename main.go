package main

import (
	"fmt"
	"net/http"
)

func main() {
	name := readName()
	fmt.Println("Hello,", name)

	// Set up a simple HTTP route
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", name)
	})

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
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
