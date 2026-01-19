package main

import (
	"fmt"
	"time"

	"go_first_ever/user"
)

func main() {
	fmt.Println("=======================================")
	fmt.Println("User App with Age Calculation===========")
	fmt.Println("=======================================\n")

	// Collect user details
	var name string
	var dob time.Time

	readName(&name)
	readDOB(&dob)

	// Calculate age
	age := calculateAge(dob)

	// Create a User object
	u := user.User{
		Name: name,
		Role: "Member", // default role
	}

	fmt.Println("\n--- User Info ---")
	fmt.Println("Name:", u.Name)
	fmt.Println("Date of Birth:", dob.Format("2006-01-02"))
	fmt.Println("Age:", age)
	fmt.Println("Role:", u.GetRole())

	// Demonstrate login
	if u.Login() {
		fmt.Println(u.Name, "successfully logged in.")
	}

	// Optional: create an Admin
	admin := user.Admin{
		User:       user.User{Name: "Alice", Role: "Admin"},
		Permission: "Full Access",
	}

	fmt.Println("\n--- Admin Info ---")
	fmt.Println("Name:", admin.Name)
	fmt.Println("Role:", admin.GetRole())
	fmt.Println("Permissions:", admin.GetPermissions())

	admin.ManageUsers()
	if admin.CanAccess("Server") {
		fmt.Println(admin.Name, "can access the Server resource.")
	}
}

// --------------------
// Input helpers

func readName(name *string) {
	for {
		fmt.Print("What is your name: ")
		_, err := fmt.Scan(name)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid name.")
			continue
		}
		if *name == "" {
			fmt.Println("Name cannot be empty.")
			continue
		}
		break
	}
}

func readDOB(dob *time.Time) {
	for {
		fmt.Print("Enter your date of birth (YYYY-MM-DD): ")
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Invalid input.")
			clearInput()
			continue
		}

		parsedDOB, err := time.Parse("2006-01-02", input)
		if err != nil {
			fmt.Println("Invalid date format. Use YYYY-MM-DD.")
			continue
		}

		if parsedDOB.After(time.Now()) {
			fmt.Println("Date of birth cannot be in the future.")
			continue
		}

		*dob = parsedDOB
		break
	}
}

func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

func clearInput() {
	var discard string
	fmt.Scanln(&discard)
}
