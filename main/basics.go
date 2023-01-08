package main

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}
type SimpleUser struct {
	FirstName string
	LastName  string
}

func oldMain() {
	var stringToDeclare = "Name"
	var whatWasSaid = saySomething()
	var whatWasSaidTwice, theOther = saySomethingTwice()
	fmt.Println(`something to print `, stringToDeclare, whatWasSaid, whatWasSaidTwice, theOther)
	pointerChange(&whatWasSaid)
	fmt.Println(`something to print `, stringToDeclare, whatWasSaid, whatWasSaidTwice, theOther)
	user := User{
		"Nome",
		"Cognome",
		"123",
		12,
		time.Now(),
	}
	log.Println("Trying type: ", user.FirstName, user)
	user.printFirstName()
}

// generic function
func saySomething() string {
	return "Something"
}

// double type return
func saySomethingTwice() (string, string) {
	return "Something", "Else"
}

// pointers
func pointerChange(s *string) {
	newValue := "Value changed"
	*s = newValue
}

// struct receiver
func (s *User) printFirstName() {
	log.Println(s.FirstName)
}
