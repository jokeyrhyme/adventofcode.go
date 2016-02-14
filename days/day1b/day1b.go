package day1b

import (
	"fmt"
	"strings"
)

// Go : main entry point
func Go(input string) {
	fmt.Println("day 1b")

	reader := strings.NewReader(input)
	var (
		ch  rune
		err error
	)
	floor := 0
	position := 0
	for floor > -1 && err == nil {
		ch, _, err = reader.ReadRune()
		if string(ch) == "(" {
			floor++
		}
		if string(ch) == ")" {
			floor--
		}
		position++
	}

	fmt.Printf("answer: %d\n", position)
}
