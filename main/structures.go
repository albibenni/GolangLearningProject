package main

import (
	"log"
	"sort"
)

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

	mySlice := makeASlice(me)
	log.Println(mySlice)

	numbers := []int{10, 2, 8, 4, 5}
	sort.Ints(numbers)
	log.Println(numbers)
	log.Println(numbers[0:3])
}

func makeASlice(me SimpleUser) []SimpleUser {
	var mySlice []SimpleUser

	mySlice = append(mySlice, me)
	return mySlice
}
