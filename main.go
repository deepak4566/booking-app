package main

import (
	"fmt"
	"strings"
)

var conferenceName string = "Goconference"

const conferenceTickets uint = 50

var remainingTickets uint = 50

func main() {

	var firstName string
	var lastName string
	var email string
	var TicketsByUser uint

	for {

		fmt.Printf("welcome to our %v and total tickets in conference is %v \n", conferenceName, conferenceTickets)

		fmt.Println("Please enter your first name:")
		fmt.Scanln(&firstName)

		fmt.Println("Please enter your last name:")
		fmt.Scanln(&lastName)

		fmt.Println("Please enter your email:")
		fmt.Scanln(&email)

		fmt.Println("Please enter no of tickets:")
		fmt.Scanln(&TicketsByUser)

		remainingTickets = remainingTickets - TicketsByUser

		if remainingTickets == 0 {
			fmt.Println("No more tickets available. The conference is sold out.")
			break
		}

		ifFirstName := len(firstName) > 2
		ifLastName := len(lastName) > 2
		ifEmail := strings.Contains(email, "@")

		ifTicketsCond := TicketsByUser <= remainingTickets

		if ifFirstName && ifLastName && ifEmail && ifTicketsCond {
			fmt.Printf("Hello %v %v, you have successfully registered for the conference\n", firstName, lastName)

		} else {
			if !ifFirstName || !ifLastName {
				fmt.Println("Please enter valid name with more than 2 characters")
			}
			if !ifEmail {
				fmt.Println("Please enter valid email")
			}

			if !ifTicketsCond {
				fmt.Println("Please enter valid no of  tickets")
			}
		}

		fmt.Printf("remaining tickets after is %v\n", remainingTickets)
	}

}
