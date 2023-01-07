package main

import "fmt"

func main() {
	var stringToDeclare = "Name"
	var whatWasSaid = saySomething()
	fmt.Println(`something to print `, stringToDeclare, whatWasSaid)
}

func saySomething() string {
	return "Something"
}
