package main

import (
	"fmt"
	"strings"

	//"sync"
	"time"
)

// Package variables, we dont have to pass these values     Scope : Package Level
// everytime, as these are main class level valuea.
var conferenceName = "Go conference"

const conferenceTickets = 50

var remainingTickets uint = 50

//var bookings = make([]map[string]string, 0)   // here is how you make a empty map
var bookings = make([]UserData, 0) //we made a empty list of objects from struct "UserData"

//structures can be compared to classes in java, it's like a light-weight class (doesnt support inheritance)
type UserData struct {
	firstName  string
	secondName string
	email      string
	userTicket int
}

// var wg = sync.WaitGroup{}

func main() {

	// greeting function calls this method here.
	greetUser()

	for {

		// take the user input from this section and returns the user input for validation
		firstName, email, userTickets, secondName := getUserInput()

		// validates the user input here and returns multiple values for multiple inputs
		isValidName, ValidEmail, isValidNumber := ValidatesInput(firstName, secondName, email, userTickets)

		if isValidName && ValidEmail && isValidNumber {

			//the main logic here
			bookTicket(userTickets, firstName, secondName, email)

			//wg.Add(1)
			go sendTicket(userTickets, firstName, secondName, email) //starts a new routine, multi-threading achcived

			//iteretes through the booking slice and amke another slice with only first name in it
			firstName := getFirstNames()

			fmt.Printf("These are all our bookings %v \n", firstName)

			if remainingTickets == 0 {
				fmt.Println("House-FULL")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First or Last name might be invalid.try again")
			}
			if !ValidEmail {
				fmt.Println("Invalid email address")
			}
			if !isValidNumber {
				fmt.Println("Number of tickets are invalid")
			}
		}
		// wg.Wait()
	}

}

func greetUser() {

	fmt.Printf("Welcome to our %v booking application", conferenceName)
	fmt.Println()
	fmt.Printf("We have total of %v tickets and %v are still remaining, Sale is on!", conferenceTickets, remainingTickets)
	fmt.Println()
	fmt.Println("Get your tickets now!!")

}
func bookTicket(userTickets int, firstName string, secondName string, email string) {
	remainingTickets = remainingTickets - uint(userTickets)

	//creating a map
	//var userData = make(map[string]string)

	var userData = UserData{
		firstName:  firstName,
		secondName: secondName, // this is like a object of our struct,which we initiating here
		email:      email,
		userTicket: userTickets,
	}
	/*userData["firstName"] = firstName
	userData["secondName"] = secondName
	userData["email"] = email
	userData["userTicket"] = strconv.FormatInt(int64(userTickets), 10)*/

	bookings = append(bookings, userData)

	fmt.Println("Thank you for buying tickets, We hope you'll enjoy your show")
	fmt.Printf("User %v has bought %v tickets for %v \n", firstName, userTickets, conferenceName)
	fmt.Printf("The tickits were send to you email @%v", email)
	fmt.Println()
	fmt.Println(time.Now())
	fmt.Printf("Now, only %v tickets are remaining \n", remainingTickets)
	fmt.Println()
	fmt.Printf("Booking Details are %v \n", bookings)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, element := range bookings {
		firstNames = append(firstNames, element.firstName) // here using/accessing the values from the "userData" object of "UserData" struct
	}

	return firstNames
}

func ValidatesInput(firstName string, secondName string, email string, userTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(secondName) >= 2
	ValidEmail := strings.Contains(email, "@")
	isValidNumber := userTickets <= int(remainingTickets) && userTickets >= 0

	return isValidName, ValidEmail, isValidNumber
}

func getUserInput() (string, string, int, string) {
	var firstName string
	var email string
	var userTickets int
	var secondName string

	//taking the inputs here
	fmt.Println("Please enter your firstName")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your secondName")
	fmt.Scan(&secondName)

	fmt.Println("Please enter your email address")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets your want.")
	fmt.Scan(&userTickets)

	return firstName, email, userTickets, secondName
}

func sendTicket(userTickets int, firstName string, secondName string, email string) {
	time.Sleep(10 * time.Second) // stops the current thread(goroutine) execution for defined duration
	var ticket = fmt.Sprintf("%v tickets for %v %v \n", userTickets, firstName, secondName)
	fmt.Println("---------------------------")
	fmt.Printf("Sending the ticket: \n %v to email address %v \n", ticket, email)
	fmt.Println("---------------------------")
	//wg.Done()
}
