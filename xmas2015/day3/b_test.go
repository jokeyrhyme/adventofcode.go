package day3

import (
	"io/ioutil"
	"testing"
)

func TestBExample1(t *testing.T) {
	input := "^v"
	expected := 3
	actual := answerB(input)
	if expected != actual {
		t.Fail()
	}
}

func TestBExample2(t *testing.T) {
	input := "^>v<"
	expected := 3
	actual := answerB(input)
	if expected != actual {
		t.Fail()
	}
}

func TestBExample3(t *testing.T) {
	input := "^v^v^v^v^v"
	expected := 11
	actual := answerB(input)
	if expected != actual {
		t.Fail()
	}
}

func TestBInput(t *testing.T) {
	bytes, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		t.Fail()
	}
	input := string(bytes[:])
	result := answerB(input)
	t.Logf("day 3: A: %d", result)
	if result != 2639 {
		t.Fail()
	}
}
