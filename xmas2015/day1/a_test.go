package day1

import (
	"io/ioutil"
	"testing"
)

func TestAInput(t *testing.T) {
	bytes, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		t.Fail()
	}
	input := string(bytes[:])
	result := answerA(input)
	t.Logf("day 1: A: %d", result)
}
