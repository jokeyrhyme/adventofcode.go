package day3

import (
	"testing"
)

func TestAEast(t *testing.T) {
	input := ">"
	expected := 2
	actual := answerA(input)
	if expected != actual {
		t.Fail()
	}
}

func TestANorthEastSouthWest(t *testing.T) {
	input := "^>v<"
	expected := 4
	actual := answerA(input)
	if expected != actual {
		t.Fail()
	}
}

func TestANorthSouthNorthSouth(t *testing.T) {
	input := "^v^v^v^v^v"
	expected := 2
	actual := answerA(input)
	if expected != actual {
		t.Fail()
	}
}
