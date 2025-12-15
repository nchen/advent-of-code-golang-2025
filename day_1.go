package main

import (
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

func norm(n int) int {
	n %= 100
	if n < 0 {
		n += 100
	}
	return n
}

func main() {
	input, err := os.ReadFile("inputs/day1.txt")
	check(err)

	lines := strings.Split(string(input), "\n")

	part_1(lines)
	part_2(lines)
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

		pos = norm(pos)

		if pos == 0 {
			zeros++
		}
	}
	fmt.Printf("Zeros encountered: %d\n", zeros)
}

func part_2(lines []string) {
	fmt.Println("Part 2")
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

		passes := 0
		stepsToZero := 0

		if dir == 'R' {
			if pos == 0 {
				stepsToZero = 100
			} else {
				stepsToZero = 100 - pos
			}
		} else { // L
			if pos == 0 {
				stepsToZero = 100
			} else {
				stepsToZero = pos
			}
		}

		if steps >= stepsToZero {
			passes = (steps-stepsToZero)/100 + 1
		}

		zeros += passes

		if dir == 'R' {
			pos += steps
		} else { // L
			pos -= steps
		}

		pos = norm(pos)
	}
	fmt.Printf("Zeros encountered: %d\n", zeros)
}
