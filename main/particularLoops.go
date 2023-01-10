package main

import "log"

func main3() {
	testForRangeLoop()
	log.Println()
	testForRangeLoopUnderscore()
	log.Println()
	log.Println()
	testForRangeLoopUnderscoreWMap()
	log.Println()
	testForRangeLoopWMap()
}

func testForRangeLoop() {
	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	for i, animal := range animals {
		log.Println(i, animal)
	}
}

// _ you dont care about the variable
func testForRangeLoopUnderscore() {
	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	for _, animal := range animals {
		log.Println(animal)
	}
}

func testForRangeLoopUnderscoreWMap() {
	animals := make(map[string]string)

	animals["dog"] = "doggyName"
	animals["cat"] = "catName"

	for _, animal := range animals {
		log.Println(animal)
	}
}
func testForRangeLoopWMap() {
	animals := make(map[string]string)

	animals["dog"] = "doggyName"
	animals["cat"] = "catName"

	for animalKind, animal := range animals {
		log.Println(animalKind, animal)
	}
}
