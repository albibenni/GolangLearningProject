package main

import "fmt"

func main() {
	var stringToDeclare = "Name"
	var whatWasSaid = saySomething()
	var whatWasSaidTwice, theOther = saySomethingTwice()
	fmt.Println(`something to print `, stringToDeclare, whatWasSaid, whatWasSaidTwice, theOther)
	pointerChange(&whatWasSaid)
	fmt.Println(`something to print `, stringToDeclare, whatWasSaid, whatWasSaidTwice, theOther)
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
