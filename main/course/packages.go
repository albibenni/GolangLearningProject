package course

import (
	"golangLearningProject/helpers"
	"log"
)

func mainIsBack() {
	var myVar helpers.SomeType
	myVar.TypeName = "Name"
	log.Println(myVar.TypeName)
}
