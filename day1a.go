package main

import (
	"fmt"
	"strings"
)

// Day1a : output solution
func Day1a() {
	fmt.Println("day 1a")

	reader := strings.NewReader(getInput(1))
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
