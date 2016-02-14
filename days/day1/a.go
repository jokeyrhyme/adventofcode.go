package day1

import (
	"fmt"
	"strings"
)

// GoA : main entry point
func GoA(input string) {
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
