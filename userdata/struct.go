package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	firstName := getUserData("Please Enter your firstName: ")
	lastName := getUserData("Please Enter your lastName: ")
	birthDate := getUserData("Please Enter your birthdate (MM/DD/YYYY): ")

	appUser := user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}

	outputUserDetails(appUser)
}

func outputUserDetails(u user) {
	fmt.Println("User Details:")
	fmt.Printf("First Name: %s\n", u.firstName)
	fmt.Printf("Last Name: %s\n", u.lastName)
	fmt.Printf("Birth Date: %s\n", u.birthDate)
	fmt.Printf("Created At: %s\n", u.createdAt.Format(time.RFC1123))
}

func getUserData(promptText string) string {
	fmt.Println(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
