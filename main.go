package main

import (
	"fmt"
	"go-learning-go/helper"
	"strconv"
)

const ZERO uint = 0
const conferenceName string = "Go Conference"

var remainingTickets uint = 50

// var bookings []string <-- Slice of strings
// init the map
var bookings = make([]map[string]string, 10) // -> Array of maps

func main() {
	greetUsers()
	for {
		//ask user their name
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTickets := helper.ValidUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTickets {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			// end program
			if remainingTickets == ZERO {
				fmt.Println("Our conference is booked out. Come back next year ")
				break // end program
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered does not contains '@' ")
			}
			if !isValidTickets {
				fmt.Println("NUmber of tickets you entered is invalid")
			}
		}
	}
}

func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func getFirstNames() (firstNames []string) {
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Get your tickets here to attend\n")
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	// update remaining tickets
	remainingTickets -= userTickets

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	bookings = append(bookings, userData)
	fmt.Printf("Name %v %v and want %v via email: %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remainning ", remainingTickets)
}

//fmt.Printf("The whole Array:    %v \n", bookings)
//fmt.Printf("The first value:    %v \n", bookings[0])
//fmt.Printf("The Slice type:     %T \n", bookings)
//fmt.Printf("The Slice length:   %v \n", len(bookings))

/*
city := "London"
	switch city {
	case "New York":
	// execute something on New York
	case "Singapore", "Hong Kong":
	case "London", "Berlin": // "London" or "Berlin"
	case "Mexico city":
	default:
		fmt.Println("No valid city selected")
	}
*/
