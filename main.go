package main

import (
	"fmt"
	"sync"
	"time"
)


var ConferenceName string = "Go Conference"
const ConferenceTickets int = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)


type UserData struct{
	firstName string
	lastName string
	email string
	UserTickets uint
}

var wg = sync.WaitGroup{}
 
func main() {


	greetings()
	firstName, lastName, email, UserTickets := UserInput()

	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInputs(firstName, lastName, email, UserTickets, remainingTickets)

	if isValidEmail && isValidName && isValidTicketNumber{
		BookTickets(UserTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(UserTickets, firstName, lastName, email)
		//printing first names of users
			
		printnames := printNames()
		fmt.Printf("First names of bookings: %v \n", printnames)

		
		if remainingTickets == 0{
			fmt.Println("Our conference is fully booked. Waiting for you next year!")
			
		}else{
			if !isValidName{
				fmt.Println("Your first name or second name is invalid!")
			}
			if !isValidEmail{
				fmt.Println("Your email is invalid!")
			}
			if !isValidTicketNumber{
				fmt.Printf("We have only %v tickets. You are trying to book %v tickets!\n", remainingTickets, UserTickets)
			}
		}
		wg.Wait() 
	}
	}



func greetings(){
	fmt.Printf("Welcome to %v booking application\n", ConferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", ConferenceTickets, remainingTickets )
	fmt.Printf("Get your tickets here to attend\n")
}

func printNames() []string{
	firstnames := []string{}
	for _, booking := range bookings{
		firstnames = append(firstnames, booking.firstName) 
	}
	return firstnames
}

func UserInput() (string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var UserTickets uint
	
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	
	fmt.Println("Enter your last email: ")
	fmt.Scan(&email)
		
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&UserTickets)
	return firstName, lastName, email, UserTickets
}

func BookTickets(UserTickets uint,firstName string, lastName string, email string){
	remainingTickets = remainingTickets	- UserTickets
	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		UserTickets: UserTickets,
	}
	  

	bookings = append(bookings, userData) 
	fmt.Printf("List of bookings: %v \n", bookings)
	
		
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation on your email: %v\n", firstName, lastName, UserTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, ConferenceName)
}

func sendTicket(UserTickets uint, firstName string, lastName string, email string){
	time.Sleep(5* time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v \n", UserTickets, firstName, lastName)
	fmt.Println("########################")
	fmt.Printf("Sending %v to %v email adress\n", ticket, email)
	fmt.Println("########################")
	wg.Done()
}