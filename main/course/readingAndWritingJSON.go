package course

import (
	"encoding/json"
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
}
