package main

import (
	"fmt"
	"os"
	"strconv"
)

func getInput(day int) string {
	data, err := Asset("data/day" + strconv.Itoa(day) + "input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(data[:])
}
