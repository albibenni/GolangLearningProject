package main

import (
	"golangLearningProject/helpers"
	"log"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Name"
	log.Println(myVar.TypeName)
}
