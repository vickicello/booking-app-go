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

	// Create an array of strings
	// Arrays in Go have a fixed size (50 elements in our case)
	// example: var bookings = [50]string

	// Slices are an abstraction of an Array in Go
	// They allow for variable length arrays or can have a sub-array
	// Also index-based but will resize if needed (note how no size is given)
	var bookings []string

	// If setting a var inline, Go can infer the datatype ("" is a string, 5 is an integer)
	// If you're not going to set it right away, you'll need to declare the dataType:
	// Like `var firstName string` or `var userTickets int`
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// get user input
	fmt.Println("What is your first name?")
	// Need to add the pointer (&) to show where the firstName value lives in memory
	// variable is firstName
	// pointer is &firstName - another var that points to the memory address of the firstName var
	// pointers are used in C, C++ - they might be used in other languages but aren't exposed to us as devs

	// So now the Scan function can grab the user input and assign it to the firstName var in memory
	fmt.Scan(&firstName)
	fmt.Println("What is your last name?")
	fmt.Scan(&lastName)

	fmt.Println("What is your email?")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you want?")
	fmt.Scan(&userTickets)
	fmt.Printf("User %v %v bought %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("You will receive a verification email at %v.\n", email)
	remainingTickets = conferenceTickets - userTickets

	// Add each user that booked to the slice - will select the next available index
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("Just this user: %v\n", bookings[0])

	// Look at the type and size of the slice
	fmt.Printf("Type: %T\n", bookings)
	fmt.Printf("Length: %v\n", len(bookings))
	fmt.Printf("There are now %v tickets remaining\n", remainingTickets)
}

// to execute the program:
// go run main.go

// this tutorial courtesy of
// Techworld with Nana!
