// Everything in Go is organized into Packages
// package main in your main.go file

package main

import "fmt"

// define the entrypoint - needs to be func main
func main() {
	// declare a variable (vars can be reassigned, consts can't)
	var conferenceName = "Go Conference"
	// constants use const
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	// %v = placeholder value
	fmt.Printf("Welcome to the %v ticketing application!\n", conferenceName)
	// string interpolation
	fmt.Println("Get your tickets for the", conferenceName, "here.")
	fmt.Println("We have", remainingTickets, "tickets left.")

	// If setting a var inline, Go can infer the datatype ("" is a string, 5 is an integer)
	// If you're not going to set it right away, you'll need to declare the dataType:
	// Like `var userName string` or `var userTickets int`
	var userName string
	var userTickets uint

	// get user input
	fmt.Println("What is your username?")
	fmt.Scan(&userName)

	fmt.Println("How many tickets do you want?")
	fmt.Scan(&userTickets)
	fmt.Printf("User %v bought %v tickets.\n", userName, userTickets)

	remainingTickets = conferenceTickets - userTickets
	fmt.Printf("There are now %v tickets remaining\n", remainingTickets)
}

// to execute the program:
// go run main.go

// this tutorial courtesy of
// Techworld with Nana!
