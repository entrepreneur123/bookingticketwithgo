package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets int  = 50
	var remainingTickets uint = 50
	var bookings []string
	
	fmt.Printf("welcome to %v conference booking app\n", conferenceName);
	fmt.Printf("we have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here")
	for
	{
		var firstName string
		var lastName string
		var email string
		var userTickets uint

	//ask user for their name
		fmt.Println("Enter your firstname")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email")
		fmt.Scan(&email)

		fmt.Println("Enter ticket you wanna to purchase")
		fmt.Scan(&userTickets)
		remainingTickets = remainingTickets -(userTickets)
		bookings = append(bookings, firstName + " "+ lastName)

		fmt.Printf("the whole slice: %v\n", bookings)
		fmt.Printf("the first value: %v\n", bookings[0])
		fmt.Printf("the slice type : %T\n", bookings)
		fmt.Printf("the slice length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v has purchased %v ticket.\n", firstName,lastName,userTickets);
		fmt.Printf("%v ticket is remaining for %v\n", remainingTickets, conferenceName);
		fmt.Printf("these are all our bookings: %v\n", bookings)


	}

	


}