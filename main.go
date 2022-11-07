package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceName = "Go Lang"
const conferenceTickets uint = 50

var remainingTickets uint = conferenceTickets
var bookings = make([]UserData, 0)

var wg = sync.WaitGroup{}

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

func main() {

	greetUser()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isNameValid, isEmailValid, isValidInputTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isNameValid && isEmailValid && isValidInputTickets {

			bookTicket(firstName, lastName, userTickets, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

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
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to Our %v Conference Booking App.\n", conferenceName)
	fmt.Printf("We have total %v tickets & %v are remaining!\n", conferenceTickets, remainingTickets)
	fmt.Println("👇👇 Grab your ticket Now!! 👇👇")
}

func printFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("Our Bookings: %v\n", firstNames)
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

	// Create a Map for User

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. We will send you confirmation @ %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("We have %v tickets remaining for the %v conference.\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#######################################################")
	fmt.Printf("Sending ticket:\n%v to %v email address\n", ticket, email)
	fmt.Println("#######################################################")

	wg.Done() 
}
