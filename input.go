package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(day int) string {
	data, err := Asset("data/day" + strconv.Itoa(day) + "input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(data[:]))
}
