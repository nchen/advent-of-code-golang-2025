package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	input, err := os.ReadFile("inputs/day1.txt")
	check(err)

	lines := strings.Split(string(input), "\n")

	part_1(lines)
}

func part_1(lines []string) {
	fmt.Println("Part 1")
	fmt.Println("Number of lines: ", len(lines))

	pos := 50
	zeros := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		dir := line[0]
		steps, err := strconv.Atoi(line[1:])
		check(err)

		if dir == 'R' {
			pos += steps
		} else { // L
			pos -= steps
		}

		pos = pos % 100

		if pos == 0 {
			zeros++
		}
	}
	fmt.Printf("Zeros encountered: %d\n", zeros)
}
