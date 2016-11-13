package day5

import (
	"regexp"
	"strings"
	"sync"
)

var doubleLetters *regexp.Regexp // "aa|bb|cc|..."
var naughtyPair, _ = regexp.Compile("ab|cd|pq|xy")
var threeVowels, _ = regexp.Compile("[aeiou].*[aeiou].*[aeiou]")

func compileDoubleLetters() {
	pairs := make([]string, 26)
	for i := range pairs {
		ascii := byte(i + 97)
		pairs[i] = strings.Repeat(string([]byte{ascii}[:]), 2)
	}
	pattern := strings.Join(pairs, "|")
	doubleLetters, _ = regexp.Compile(pattern)
}

func isNice(word string) bool {
	compileDoubleLetters()
	bytes := []byte(word)
	if naughtyPair.Match(bytes) {
		return false
	}
	return doubleLetters.Match(bytes) && threeVowels.Match(bytes)
}

func countNiceWord(word string) int {
	if isNice(word) {
		return 1
	}
	return 0
}

func countNiceWords(done <-chan struct{}, words []string) <-chan int {
	c := make(chan int)
	go func() {
		var wg sync.WaitGroup
		for i := range words {
			wg.Add(1)
			go func(word string) {
				select {
				case c <- countNiceWord(word):
				case <-done:
				}
				wg.Done()
			}(words[i])
		}
		go func() {
			wg.Wait()
			close(c)
		}()
	}()
	return c
}

func words(input string) []string {
	return strings.Split(input, "\n")
}

func answerA(input string) int {
	done := make(chan struct{})
	defer close(done)

	c := countNiceWords(done, words(input))
	result := 0
	for r := range c {
		result += r
	}

	return result
}
