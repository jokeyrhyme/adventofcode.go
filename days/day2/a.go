// http://divan.github.io/posts/go_concurrency_visualize/
// https://blog.golang.org/pipelines

package day2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

func presents(input string) []string {
	return strings.Split(input, "\n")
}

func dimensions(present string) []int {
	result := make([]int, 3)
	for index, value := range strings.Split(present, "x") {
		num, _ := strconv.Atoi(value)
		result[index] = num
	}
	return result
}

func area(sides []int) int {
	return sides[0] * sides[1]
}

func areas(dimensions []int) []float64 {
	result := make([]float64, 3)
	result[0] = float64(dimensions[0] * dimensions[1])
	result[1] = float64(dimensions[1] * dimensions[2])
	result[2] = float64(dimensions[2] * dimensions[0])
	return result
}

func smallestArea(dimensions []int) int {
	a := areas(dimensions)
	return int(math.Min(math.Min(a[0], a[1]), a[2]))
}

func surfaceArea(dimensions []int) int {
	w := dimensions[0]
	h := dimensions[1]
	l := dimensions[2]
	return 2*l*w + 2*w*h + 2*h*l
}

func wrappingPaper(present string) int {
	dim := dimensions(present)
	return surfaceArea(dim) + smallestArea(dim)
}

func totalWrappingPaper(done <-chan struct{}, presents []string) <-chan int {
	c := make(chan int)
	go func() {
		var wg sync.WaitGroup
		for _, p := range presents {
			wg.Add(1)
			go func(present string) {
				select {
				case c <- wrappingPaper(present):
				case <-done:
				}
				wg.Done()
			}(p)
		}
		go func() { // HL
			wg.Wait()
			close(c) // HL
		}()
	}()
	return c
}

// GoA : main entry point
func GoA(input string) {
	fmt.Println("day 2a")

	done := make(chan struct{})
	defer close(done)

	c := totalWrappingPaper(done, presents(input))
	result := 0
	for r := range c {
		result += r
	}

	fmt.Printf("answer: %d\n", result)
}
