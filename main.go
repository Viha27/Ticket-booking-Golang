package main

import (
	"fmt"
	"strings"
)

func main() {
	confName := "Go Conference"
	const confTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("confName is %T type, confTickets is %T type, and remainingTickets are %T type\n", confName, confTickets, remainingTickets)

	fmt.Printf("Welcome to our %v booking app\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available", confTickets, remainingTickets)
	fmt.Println("Get your tickets to attend", confName)

	for {
		var firstName string
		var lastName string
		var email string
		//ask user for their name
		var userTickets uint
		//var bookings []string
		//var bookings = []string{}
		bookings := []string{}

		fmt.Println("Enter your first name :")
		fmt.Scan(&firstName)
		fmt.Println("Enter your Last name :")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email :")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The whole slice:%v\n", bookings)
		fmt.Printf("First Value:%v\n", bookings[0])
		fmt.Printf("Slice Type:%T\n", bookings)
		fmt.Printf("Slice length:%v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets, you will recive confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confName)

		firstNames := []string{}
		for _, value := range bookings {
			var names = strings.Fields(value)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first name of the bookings are: %v", firstNames)

	}

}
