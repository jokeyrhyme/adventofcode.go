package day4

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
	expected := 1038736
	result := answerB(input)
	t.Logf("day 4: B: %d", result)
	if result != expected {
		t.Fail()
	}
}
