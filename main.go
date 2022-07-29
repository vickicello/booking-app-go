// Everything in Go is organized into Packages
// package main in your main.go file

package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	greetUsers(conferenceName, remainingTickets)

	for {
		var bookings []string
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// get user input
		fmt.Println("What is your first name?")
		fmt.Scan(&firstName)

		fmt.Println("What is your last name?")
		fmt.Scan(&lastName)

		fmt.Println("What is your email?")
		fmt.Scan(&email)

		fmt.Println("How many tickets do you want?")
		fmt.Scan(&userTickets)

		isValidName, isValidEmail, isValidTicketCount := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// Validate name, email, and tickets requested
		if isValidName && isValidEmail && isValidTicketCount {
			remainingTickets = conferenceTickets - userTickets

			fmt.Printf("User %v %v bought %v tickets.\n", firstName, lastName, userTickets)
			fmt.Printf("You will receive a verification email at %v.\n", email)

			// Add each user that booked to the slice - will select the next available index
			bookings = append(bookings, firstName+" "+lastName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("Here are all of our bookings: %v\n", firstNames)

			fmt.Printf("There are now %v tickets remaining\n", remainingTickets)

			if remainingTickets == 0 {
				fmt.Println("Sorry, the conference is sold out!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The first name or last name you entered is too short.\n")
			}
			if !isValidEmail {
				fmt.Println("The email you entered does not have an @ sign.\n")
			}
			if !isValidTicketCount {
				fmt.Printf("Number of tickets is invalid, please try again\n")
			}
		}
	}
}

func greetUsers(confName string, remainTickets uint) {
	fmt.Printf("Welcome to the %v ticketing application!\n", confName)
	fmt.Println("Get your tickets for the", confName, "here.")
	fmt.Println("We have", remainTickets, "tickets left.")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketCount := userTickets > 0 && userTickets < remainingTickets

	// in Go, you can return n number of return values in a function
	return isValidName, isValidEmail, isValidTicketCount
}

// to execute the program:
// go run main.go

// this tutorial courtesy of
// Techworld with Nana!
