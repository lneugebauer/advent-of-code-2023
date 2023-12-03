package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func trim(s string) string {
	return strings.TrimSpace(s)
}

func toInt(s string) int {
	i, err := strconv.Atoi(trim(s))
	check(err)
	return i
}

func process(s string) int {
	slice := strings.Split(s, ":")
	colors := strings.Split(slice[1], ";")
	var red int
	var green int
	var blue int
	for _, color := range colors {
		cubes := strings.Split(color, ",")
		for _, cube := range cubes {
			values := strings.Split(cube, " ")
			amount := toInt(values[1])
			c := trim(values[2])
			if c == "red" && amount > red {
				red = amount
			}
			if c == "green" && amount > green {
				green = amount
			}
			if c == "blue" && amount > blue {
				blue = amount
			}
		}
	}

	return red * green * blue
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		round := process(scanner.Text())
		sum += round
	}
	err = scanner.Err()
	check(err)

	fmt.Print(sum)
}
