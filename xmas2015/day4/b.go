package day4

import "strings"

func answerB(input string) int {
	input = strings.TrimSpace(input)
	number := 0
	hashHead := strings.Repeat("0", 6)
	for !testNumber(input, number, hashHead) && number < 10000000 {
		number++
	}
	return number
}
