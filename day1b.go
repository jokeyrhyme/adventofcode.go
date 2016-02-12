package main

import (
	"fmt"
	"strings"
)

// Day1b : output solution
func Day1b() {
	fmt.Println("day 1b")

	reader := strings.NewReader(getInput(1))
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
