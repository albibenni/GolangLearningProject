package exercises

import (
	"bufio"
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

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
func EOneDotFour() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
