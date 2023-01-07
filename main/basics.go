package main

import "fmt"

func main() {
	var stringToDeclare = "Name"
	var whatWasSaid = saySomething()
	var whatWasSaidTwice, theOther = saySomethingTwice()
	fmt.Println(`something to print `, stringToDeclare, whatWasSaid, whatWasSaidTwice, theOther)
}

func saySomething() string {
	return "Something"
}
func saySomethingTwice() (string, string) {
	return "Something", "Else"
}
