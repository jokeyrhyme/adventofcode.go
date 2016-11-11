package day3

import (
	"io/ioutil"
	"testing"
)

func TestAExample1(t *testing.T) {
	input := ">"
	expected := 2
	actual := answerA(input)
	if expected != actual {
		t.Fail()
	}
}

func TestAExample2(t *testing.T) {
	input := "^>v<"
	expected := 4
	actual := answerA(input)
	if expected != actual {
		t.Fail()
	}
}

func TestAExample3(t *testing.T) {
	input := "^v^v^v^v^v"
	expected := 2
	actual := answerA(input)
	if expected != actual {
		t.Fail()
	}
}

func TestAInput(t *testing.T) {
	bytes, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		t.Fail()
	}
	input := string(bytes[:])
	result := answerA(input)
	t.Logf("day 3: A: %d", result)
	if result != 2565 {
		t.Fail()
	}
}
