package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Cyber"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("welcome to our  %v  conference\n", conferenceName)
	fmt.Printf("we have %v tickets and %v are still reamaining \n", conferenceTickets, remainingTickets)

	for {
		//ask user for their name
		var firstName string
		var lastName string
		var userTickets uint
		var email string

		//using pointers to capture user input
		fmt.Println("Enter your FIRST name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your Last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address:")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		//Basic logic for booking tickets
		remainingTickets = remainingTickets - userTickets

		//declaring slices
		bookings := []string{}
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("User %v booked %v tickets. You will receive confirmation email at %v \n", firstName, userTickets, email)

		fmt.Printf("%v tickets are remaining from %v \n", remainingTickets, conferenceTickets)

		//printing first name
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("the first names of bookers are %v \n", firstNames)

	}

}
