package day1

import (
	"fmt"
	"strings"
)

func answerA(input string) int {
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

	return floor
}
