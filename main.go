package main

import (
	"fmt"
)

func main() {

	const conferenceName = "Go Lang"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets
	bookings := []string{}

	fmt.Printf("Welcome to Bookings for %v conference.🌠\n", conferenceName)
	fmt.Printf("We have total %v tickets & %v are remaining!\n", conferenceTickets, remainingTickets)
	fmt.Println("👇👇 Grab your ticket Now!! 👇👇")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	for {
		fmt.Println("Enter your First Name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your Last Name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your Email:")
		fmt.Scan(&email)

		fmt.Println("Enter total number of tickets you wanna book:")
		fmt.Scan(&userTickets)

		bookings = append(bookings, firstName+" "+lastName)
		remainingTickets = remainingTickets - userTickets

		fmt.Printf("Thank you %v %v for booking %v tickets. We will send you confirmation @ %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("We have %v tickets remaining for the %v conference.\n", remainingTickets, conferenceName)
		fmt.Printf("Our Bookins: %v\n", bookings)
	}

}
