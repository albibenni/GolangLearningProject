package exercises

import (
	"fmt"
	"os"
	"strings"
)

// Page 8.
// Echo prints its command-line arguments.
func echo() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
func EOneDotOne() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
func EOneDotTwo() {
	for index, arg := range os.Args[0:] {
		fmt.Println(index, arg)
	}
}
