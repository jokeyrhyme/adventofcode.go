package day1

import (
	"fmt"
	"strings"
)

func answerB(input string) int {
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

	return position
}
