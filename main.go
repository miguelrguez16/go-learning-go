package main

import (
	"fmt"
	"strings"
)

const EMPTY string = " "
const ZERO uint = 0
const AT string = "@"

const conferenceName string = "Go Conference"
const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings []string

/*
* WELCOME TO GO CONFERENCE
 */
func main() {
	greetUsers()
	for {
		//ask user their name
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTickets := validUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTickets {

			updateBookings(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			// end program
			if remainingTickets == ZERO {
				fmt.Println("Our conference is booked out. Come back next year ")
				break
				// end program
			}
			//  } else if condition {
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

func validUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, AT)
	var isValidTickets = userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTickets
}

func getFirstNames() (firstNames []string) {
	for _, b := range bookings {
		var names = strings.Fields(b)
		firstNames = append(firstNames, names[ZERO])
	}
	return firstNames
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have totoal of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
}

func updateBookings(userTickets uint, firstName string, lastName string, email string) {
	// update remaining tickets
	remainingTickets -= userTickets
	bookings = append(bookings, firstName+EMPTY+lastName)
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
