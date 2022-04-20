package main

import "fmt"

func main() {
	fmt.Print("Hello World\n")

	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("ConferenceName is %T\n", conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are stil available\n", conferenceTickets, remainingTickets)

	// Array
	var bookings [50]string

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask for their name
	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Printf("Enter your email: ")
	fmt.Scan(&email)
	fmt.Printf("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	remainingTickets = remainingTickets - userTickets
	bookings[0] = firstName + " " + lastName

	userTickets = 2
	fmt.Printf("Thank you %v %v for bookiong %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remainting for %v\n", remainingTickets, conferenceName)
}
