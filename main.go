// Everything in Go is organized into Packages
// package main in your main.go file

package main

import (
	"fmt"
	"time"
)

// package level variables are accessible within all functions in the package
// Although best practice is to make variables as local as possible
const conferenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// use make() to initialize a new UserData struct of size 0:
var bookings = make([]UserData, 0)

// Create a custom type of data in our app; struct is named UserData
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketCount := validateUserInput(firstName, lastName, email, userTickets)

		// Validate name, email, and tickets requested
		if isValidName && isValidEmail && isValidTicketCount {
			bookTicket(userTickets, firstName, lastName, email)
			// turn sendTicket into a goroutine, making our app concurrent
			// sendTicket will spin off into a new thread so the application won't be blocked and
			// more users can book tickets!
			go sendTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("Here are all of our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sorry, the conference is sold out!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The first name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("The email you entered does not have an @ sign.")
			}
			if !isValidTicketCount {
				fmt.Printf("Number of tickets is invalid, please try again")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to the %v ticketing application!\n", conferenceName)
	fmt.Println("Get your tickets for the", conferenceName, "here.")
	fmt.Printf("There are %v tickets for the conference, and %v are still available\n", conferenceTickets, remainingTickets)
}

// Inside the () are the function input params; outside are the output params or return values
func getFirstNames() []string {
	firstNames := []string{}
	// now we will iterate over a list of structs rather than an array of maps:
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = conferenceTickets - userTickets

	// create a struct for each user's data
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// Add each user that booked to the slice - will select the next available index
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("User %v %v bought %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("You will receive a verification email at %v.\n", email)
	fmt.Printf("There are now %v tickets remaining\n\n", remainingTickets)
}

// simulate sending ticket to a user's email:
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// simulate email sending by adding 10 seconds of sleep and blocking execution of the program
	time.Sleep(10 * time.Second)
	// Formats a string and save into var
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###############")
}

// to execute the program:
// go run main.go helper.go or go run .

// this tutorial courtesy of
// Techworld with Nana!
