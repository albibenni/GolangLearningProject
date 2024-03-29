package course

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func JsonOperations() {
	myJson := `
[
	{
		"first_name": "Clark",
		"last_name": "Kent",
		"hair_color": "Black",
		"has_dog": true
	},
	{
		"first_name": "Bruce",
		"last_name": "Wayne",
		"hair_color": "Black",
		"has_dog": false
	}
]`
	var unmashalled []Person

	err := json.Unmarshal([]byte(myJson), &unmashalled)

	if err != nil {
		log.Println("Error unmarshalling json", err)
	}
	log.Printf("Unmarshed: %v", unmashalled)

	//	Write a json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)
	var m2 Person
	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	m2.HairColor = "black"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", " ")
	if err != nil {
		log.Println("Error Marshalling: ", err)
	}
	fmt.Println(string(newJson))
}
