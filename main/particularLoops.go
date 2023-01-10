package main

import "log"

func main() {

	testForRangeLoop()
}

func testForRangeLoop() {
	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	for i, animal := range animals {
		log.Println(i, animal)
	}
}
