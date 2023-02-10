// declare that this file belongs to the helper package
package main

import "strings"

// we can modularize our code by making a helper file, but we will have to include the file in our go run command:
// go run main.go helper.go
// alternatively, just specify a folder location:
// go run .

// multiple output values can be placed in () after the input params ()
// Capitalizing the first level of a function will export it to be used in main.go
// You can also capitalize the first letter of a var to export it
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketCount := userTickets > 0 && userTickets < remainingTickets

	// in Go, you can return n number of return values in a function
	return isValidName, isValidEmail, isValidTicketCount
}
