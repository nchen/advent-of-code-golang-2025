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

type Range struct {
	start int
	end   int
}

func scanRanges(lines []string) []Range {
	var ranges []Range

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			continue
		}

		start, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		end, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err1 != nil || err2 != nil {
			continue
		}

		ranges = append(ranges, Range{start: start, end: end})
	}
	return ranges
}

func scanIDs(lines []string) []int {
	var ids []int

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 || strings.Index(line, "-") >= 0 {
			continue
		}

		id, err := strconv.Atoi(line)
		if err != nil {
			continue
		}

		ids = append(ids, id)
	}
	return ids
}

func main() {
	input, err := os.ReadFile("inputs/day5.txt")
	check(err)

	lines := strings.Split(string(input), "\n")

	part_1(lines)
}

func part_1(lines []string) {
	fmt.Println("Part 1")
	fmt.Println("Number of lines: ", len(lines))

	ranges := scanRanges(lines)
	ids := scanIDs(lines)

	fmt.Println("Number of ranges: ", len(ranges))
	fmt.Println(ranges[2].start, ranges[2].end)

	fmt.Println("Number of IDs: ", len(ids))
	fmt.Println(ids[5])
}
