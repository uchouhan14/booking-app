package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	remainingTickets := 50

	fmt.Printf("Welcome to %v booking application we still have %v tickets left out of %v tickets. \n",
		conferenceName, remainingTickets, conferenceTickets)

	fmt.Println("Get your tickets from here to attend.")

	var firstName string
	var lastName string
	var userEmail string
	var userTickets int

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email")
	fmt.Scan(&userEmail)
	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)
	fmt.Printf("Thank you %v for purchasing %v tickets \n", firstName, userTickets)

	var bookings []string
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Current User List %s \n", bookings)
}

// initial release of the app
