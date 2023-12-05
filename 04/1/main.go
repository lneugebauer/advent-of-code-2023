package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func process(s string) int {
	contents := strings.Split(s, ":")[1]
	numbers := strings.Split(contents, "|")
	winningNumbers := strings.Split(numbers[0], " ")
	numbersIHave := strings.Split(numbers[1], " ")
	var points int
	for _, v := range winningNumbers {
		if string(v) == "" {
			continue
		}
		winning := slices.Contains(numbersIHave, string(v))
		if winning {
			if points == 0 {
				points = 1
			} else {
				points = points + points
			}
		}
	}
	return points
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		points := process(scanner.Text())
		sum += points
	}
	err = scanner.Err()
	check(err)

	fmt.Print(sum)
}
