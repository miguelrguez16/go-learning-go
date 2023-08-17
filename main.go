package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50
const EMPTY string = " "
const ZERO uint = 0
const conferenceName string = "Go Conference"

var remainingTickets uint = 50

/*
*
WELCOME TO GO CONFERENCE
*/
func main() {
	/* bookings
	* var bookings = [50]string ["one","two"] <-- Array
	* var bookings [50]string <-- Array
	* var bookings []string <-- Slice
	 */
	// bookings := []string{}
	var bookings []string
	var userTickets uint
	var firstName, lastName, email string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have totoal of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
	for {
		//ask user their name
		fmt.Println("Enter your first name: ")
		_, err := fmt.Scan(&firstName)
		if err != nil {
			return
		}

		fmt.Println("Enter your last name: ")
		_, err = fmt.Scan(&lastName)
		if err != nil {
			return
		}
		fmt.Println("Enter your email: ")
		_, err = fmt.Scan(&email)
		if err != nil {
			return
		}
		fmt.Println("Enter number of tickets:")
		_, err = fmt.Scan(&userTickets)
		if err != nil {
			return
		}

		if userTickets <= remainingTickets {
			// update remaining tickets
			remainingTickets -= userTickets
			// bookings[0] = firstName + EMPTY + lastName <-- Arrays
			bookings = append(bookings, firstName+EMPTY+lastName)

			fmt.Printf("Name %v %v and want %v via email: %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remainning ", remainingTickets)

			fmt.Printf("These are all our bookings: %v\n", bookings)
			// firstNames := []string{}
			var firstNames []string
			for _, b := range bookings {
				var names = strings.Fields(b)
				firstNames = append(firstNames, names[ZERO])
			}
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			// end program
			if remainingTickets == ZERO {
				fmt.Println("Our conference is booked out. Come back next year ")
				break // end loop
				// end program
			}

		} else {
			fmt.Printf("We only have %v tickets remaining, so you can not book %v tickets\n", remainingTickets, userTickets)
			// skip current execution
			continue
			// break
		}
	}

}

//fmt.Printf("The whole Array:    %v \n", bookings)
//fmt.Printf("The first value:    %v \n", bookings[0])
//fmt.Printf("The Slice type:     %T \n", bookings)
//fmt.Printf("The Slice length:   %v \n", len(bookings))
