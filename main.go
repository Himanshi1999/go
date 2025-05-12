package main

import (
	"fmt"
	"strings"

	"github.com/Himanshi1999/go.git/helper"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings []string

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTickets {
			bookTicket(firstName, lastName, email, userTickets)
			firstNames := getFirstNames()
			fmt.Printf("The firstNames of bookings are %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign.")
			}
			if !isValidTickets {
				fmt.Println("number of tickets you entered is invalid.")
			}
			continue
		}

	}
}

func greetUsers() {
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var Names = strings.Fields(booking)
		firstNames = append(firstNames, Names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
