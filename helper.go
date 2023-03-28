package main

import "strings"

func ValidateUserInputs(firstName string, lastName string, email string, UserTickets uint, remainingTickets uint) (bool, bool, bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := UserTickets > 0 && UserTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
 
}