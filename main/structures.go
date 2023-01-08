package main

import "log"

func main() {
	log.Println("Something")

	me := SimpleUser{
		FirstName: "Alb",
		LastName:  "Ben",
	}
	myMap := make(map[string]string)
	secondMap := make(map[string]SimpleUser)

	myMap["key"] = "Value"
	secondMap["me"] = me

	log.Println(myMap["key"])
	log.Println(secondMap["me"].FirstName)
}
