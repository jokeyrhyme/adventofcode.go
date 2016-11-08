package day3

import "strings"

func answerA(input string) int {
	reader := strings.NewReader(input)

	houses := make(map[string]bool)
	ns := 0
	ew := 0

	coords := string(ns) + "," + string(ew)
	houses[coords] = true

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
		coords = string(ns) + "," + string(ew)
		houses[coords] = true
	}

	return len(houses)
}
