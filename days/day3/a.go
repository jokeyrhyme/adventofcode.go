package day3

import (
	"fmt"
	"strings"
)

func answerA(input string) int {
	reader := strings.NewReader(input)

	houses := make(map[string]int)
	ns := 0
	ew := 0

	houses[string(ns)+","+string(ew)] = 1

	var (
		ch  rune
		err error
	)
	for err == nil {
		ch, _, err = reader.ReadRune()
		if string(ch) == "^" {
			ns++
		}
		if string(ch) == "v" {
			ns--
		}
		if string(ch) == ">" {
			ew++
		}
		if string(ch) == "<" {
			ew--
		}
		coords := string(ns) + "," + string(ew)
		_, ok := houses[coords]
		if ok {
			houses[coords]++
		} else {
			houses[coords] = 1
		}
	}

	visited := 0
	for _, visits := range houses {
		if visits > 0 {
			visited++
		}
	}

	return visited
}

// GoA : main entry point
func GoA(input string) {
	fmt.Println("day 3a")
	fmt.Printf("answer: %d\n", answerA(input))
}
