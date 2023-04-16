package main

import (
	"fmt"
	"golangLearningProject/main/leetCode"
)

func main() {
	//exercises.EOneDotOne()
	//exercises.EOneDotTwo()
	//course.JsonOperations()
	//course.Division(10, 2)
	//exercises.EOneDotFour()
	palindrome1 := leetCode.IsPalindrome(101)
	palindrome2 := leetCode.IsPalindrome(11)
	palindrome3 := leetCode.IsPalindrome(11111111111111101)

	fmt.Printf("n1: %t \nn2: %t \nn3: %t \n", palindrome1, palindrome2, palindrome3)
}
