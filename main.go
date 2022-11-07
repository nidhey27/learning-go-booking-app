package main

import (
	"fmt"
	"strings"
)

const conferenceName = "Go Lang"
const conferenceTickets uint = 50

var remainingTickets uint = conferenceTickets
var bookings = []string{}

func main() {

	greetUser()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isNameValid, isEmailValid, isValidInputTickets := validateUserInput(firstName, lastName, email, userTickets)

		if isNameValid && isEmailValid && isValidInputTickets {

			bookTicket(firstName, lastName, userTickets, email)

			printFirstNames()

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Better luck next time. 😅")
				break
			}

		} else {
			if !isNameValid {
				fmt.Println("First Name or Last Name is too Short!👌")
			}
			if !isNameValid {
				fmt.Println("Invalid Email, Please enter the correct one!👎")
			}
			if !isValidInputTickets {
				fmt.Println("You have entered invalid Tickets count, Tey Again!😶‍🌫️")
			}
			// fmt.Printf("We only have %v tickets left. Please Try Again with Valid Input.\n", remainingTickets)
		}

	}

}

func greetUser() {
	fmt.Printf("Welcome to Our %v Conference Booking App.\n", conferenceName)
	fmt.Printf("We have total %v tickets & %v are remaining!\n", conferenceTickets, remainingTickets)
	fmt.Println("👇👇 Grab your ticket Now!! 👇👇")
}

func printFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings {
		var name = strings.Fields(booking)
		firstNames = append(firstNames, name[0])
	}
	fmt.Printf("Our Bookins: %v\n", firstNames)
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isNameValid := len(firstName) > 2 && len(lastName) > 2
	isEmailValid := strings.Contains(email, "@")
	isValidInputTickets := userTickets > 0 && userTickets <= remainingTickets

	return isNameValid, isEmailValid, isValidInputTickets
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your First Name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last Name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email:")
	fmt.Scan(&email)

	fmt.Println("Enter total number of tickets you wanna book:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, userTickets uint, email string) {
	bookings = append(bookings, firstName+" "+lastName)
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. We will send you confirmation @ %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("We have %v tickets remaining for the %v conference.\n", remainingTickets, conferenceName)
}
