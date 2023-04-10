package course

import (
	"errors"
	"log"
)

func Division(x, y float32) {
	result, err := divide(x, y)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Result of the division is: ", result)
}

func divide(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("cannot divide by 0")
	}
	result = x / y
	return result, nil
}
