package exercises

import (
	"fmt"
	"os"
	"strings"
)

// Page 6.
// Echo2 prints its command-line arguments.

func echo() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
func EOneDotOne() {
	echo()
	fmt.Println(strings.Join(os.Args[0:], " "))

}
