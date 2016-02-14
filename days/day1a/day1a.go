package day1a

import (
	"fmt"
	"strings"
)

// Go : main entry point
func Go(input string) {
	fmt.Println("day 1a")

	reader := strings.NewReader(input)
	var (
		ch  rune
		err error
	)
	floor := 0
	for err == nil {
		ch, _, err = reader.ReadRune()
		if string(ch) == "(" {
			floor++
		}
		if string(ch) == ")" {
			floor--
		}
	}

	fmt.Printf("answer: %d\n", floor)
}
