package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isNumber(r rune) bool {
	if _, err := strconv.Atoi(string(r)); err == nil {
		return true
	}

	return false
}

func analyzeString(s string) string {
	var first string
	var last string

	for _, r := range s {
		if first == "" && isNumber(r) {
			first = string(r)
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		r := rune(s[i])
		if last == "" && isNumber(r) {
			last = string(r)
		}
	}

	return first + last
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		s := analyzeString(scanner.Text())
		v, err := strconv.Atoi(s)
		check(err)
		sum += v
	}
	fmt.Print(sum)

	err = scanner.Err()
	check(err)
}
