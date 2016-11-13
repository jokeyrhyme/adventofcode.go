package day4

import (
	"io/ioutil"
	"testing"
)

func TestAExample1(t *testing.T) {
	input := "abcdef"
	expected := 609043
	actual := answerA(input)
	if expected != actual {
		t.Fail()
	}
}

func TestAExample2(t *testing.T) {
	input := "pqrstuv"
	expected := 1048970
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
	expected := 254575
	result := answerA(input)
	t.Logf("day 4: A: %d", result)
	if result != expected {
		t.Fail()
	}
}
