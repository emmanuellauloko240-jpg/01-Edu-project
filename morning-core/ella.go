package main

import (
	"fmt"
	"strings"
)

func main()  {
	

	var (
		conferenceName = "LEARN2EARN"
		conferenceTicket = 204
		remainingTicket = 204
		bookings = []string{}
	)

	fmt.Printf("WELCOME TO OUR %v CONFERENCE IN THSES CONFERENCE WE HAVE %v TICKET TO  BOOK AND %v IS REMAINING \n",conferenceName, conferenceTicket, remainingTicket)
	fmt.Println("GET YOUR OWN TICKET NOW")

	for {

	var (
		firstName string
		lastName string
		emailAddress string
		userTicker int
	)

	fmt.Println("ENTER YOUR FIRST NAME:")
	fmt.Scan(&firstName)

	fmt.Println("ENTER YOUR LAST NAME:")
	fmt.Scan(&lastName)

	fmt.Println("ENTER YOUR EMAIL ADDREESS")
	fmt.Scan(&emailAddress)

	fmt.Println("HOW MANY TICKET ARE YOU BOOKING")
	fmt.Scan(&userTicker)

	isvalidName := len(firstName) > 2 && len(lastName) > 2
	isvalidEmail := strings.Contains(emailAddress, "@gmail.com")
	isvalidTicket := userTicker > 0 && userTicker <= remainingTicket

	if isvalidName && isvalidEmail&& isvalidTicket  {

	remainingTicket = remainingTicket - userTicker
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Println("THANK YOU FOR BOOKING WITH US WE WILL SEND YOU AN EMAIL ")
	fmt.Printf("WE HAVE %v TICKET REMAINING \n",remainingTicket)
	fmt.Printf("THESE ARE ALL THE BOOKING:%v \n", bookings)

	if remainingTicket == 0 {
		fmt.Println(" OUR TICKET IS ALL BOOKED OUT")
		break
	}

	}else {
		if !isvalidName{
			fmt.Println("INVALID NAME")
		}
		if !isvalidEmail{
			fmt.Println("INVALIG EMAIL")
		}
		if !isvalidTicket{
			fmt.Println("INVALID TICKET NEEDED")
		}
	}

	}

}
	