package helper

import (
	"strings"
)

const At string = "@"

// ValidUserInput capitalized the name of the function to export it
func ValidUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, At)
	var isValidTickets = userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTickets
}
