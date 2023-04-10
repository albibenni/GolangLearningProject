package course

import "golangLearningProject/helpers"

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func channlesmain() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	print(<-intChan)
}
