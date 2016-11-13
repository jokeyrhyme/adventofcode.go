package day4

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

func md5Hash(input string) string {
	data := []byte(input)
	sum := md5.Sum(data)
	return hex.EncodeToString(sum[:])
}

func testNumber(input string, number int, hashHead string) bool {
	hash := md5Hash(input + strconv.Itoa(number))
	index := strings.Index(hash, hashHead)
	return index == 0
}

func answerA(input string) int {
	input = strings.TrimSpace(input)
	number := 0
	hashHead := strings.Repeat("0", 5)
	for !testNumber(input, number, hashHead) && number < 10000000 {
		number++
	}
	return number
}
