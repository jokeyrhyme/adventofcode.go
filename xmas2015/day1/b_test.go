package day1

import (
	"io/ioutil"
	"testing"
)

func TestBInput(t *testing.T) {
	bytes, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		t.Fail()
	}
	input := string(bytes[:])
	result := answerB(input)
	t.Logf("day 1: B: %d", result)
}
