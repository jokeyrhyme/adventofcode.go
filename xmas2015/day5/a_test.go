package day5

import (
	"io/ioutil"
	"testing"
)

func TestIsNiceExample1(t *testing.T) {
	if isNice("ugknbfddgicrmopn") != true {
		t.Fail()
	}
}

func TestIsNiceExample2(t *testing.T) {
	if isNice("aaa") != true {
		t.Fail()
	}
}

func TestIsNiceExample3(t *testing.T) {
	if isNice("jchzalrnumimnmhp") != false {
		t.Fail()
	}
}

func TestIsNiceExample4(t *testing.T) {
	if isNice("haegwjzuvuyypxyu") != false {
		t.Fail()
	}
}

func TestIsNiceExample5(t *testing.T) {
	if isNice("dvszwmarrgswjxmb") != false {
		t.Fail()
	}
}

func TestAInput(t *testing.T) {
	bytes, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		t.Fail()
	}
	input := string(bytes[:])
	expected := 255
	result := answerA(input)
	t.Logf("day 5: A: %d", result)
	if expected != result {
		t.Fail()
	}
}
