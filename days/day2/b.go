package day2

import (
	"fmt"
	"math"
	"sync"
)

func perimeter(a int, b int) int {
	return 2 * (a + b)
}

func perimeters(dimensions []int) []float64 {
	result := make([]float64, 3)
	result[0] = float64(perimeter(dimensions[0], dimensions[1]))
	result[1] = float64(perimeter(dimensions[1], dimensions[2]))
	result[2] = float64(perimeter(dimensions[2], dimensions[0]))
	return result
}

func smallestPerimeter(dimensions []int) int {
	a := perimeters(dimensions)
	return int(math.Min(math.Min(a[0], a[1]), a[2]))
}

func volume(dimensions []int) int {
	return dimensions[0] * dimensions[1] * dimensions[2]
}

func ribbon(present string) int {
	dim := dimensions(present)
	return volume(dim) + smallestPerimeter(dim)
}

func totalRibbon(done <-chan struct{}, presents []string) <-chan int {
	c := make(chan int)
	go func() {
		var wg sync.WaitGroup
		for _, p := range presents {
			wg.Add(1)
			go func(present string) {
				select {
				case c <- ribbon(present):
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

// GoB : main entry point
func GoB(input string) {
	fmt.Println("day 2b")

	done := make(chan struct{})
	defer close(done)

	c := totalRibbon(done, presents(input))
	result := 0
	for r := range c {
		result += r
	}

	fmt.Printf("answer: %d\n", result)
}
