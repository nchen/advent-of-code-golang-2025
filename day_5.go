package main

import (
	"fmt"
	"os"
	"sort"
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
	part_2(lines)
}

func part_1(lines []string) {
	fmt.Println("Part 1")
	fmt.Println("Number of lines: ", len(lines))

	ranges := scanRanges(lines)
	ids := scanIDs(lines)

	fmt.Println("Number of ranges: ", len(ranges))
	// fmt.Println(ranges[2].start, ranges[2].end)

	fmt.Println("Number of IDs: ", len(ids))
	// fmt.Println(ids[5])

	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].start < ranges[j].start {
			return true
		} else if ranges[i].start > ranges[j].start {
			return false
		} else {
			return ranges[i].end < ranges[j].end
		}
	})

	// for _, r := range ranges {
	// 	fmt.Println(r.start, "-", r.end)
	// }

	var freshIDs []int
	for _, id := range ids {
		for _, r := range ranges {
			if id < r.start {
				break
			}
			if id <= r.end {
				freshIDs = append(freshIDs, id)
				break
			}
		}
	}

	fmt.Println("Length of fresh IDs: ", len(freshIDs))
}

func part_2(lines []string) {
	fmt.Println("Part 2")
	ranges := scanRanges(lines)
	fmt.Println("Number of ranges: ", len(ranges))

	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].start < ranges[j].start {
			return true
		} else if ranges[i].start > ranges[j].start {
			return false
		} else {
			return ranges[i].end < ranges[j].end
		}
	})

	numIDs := 0
	lastStart := 0
	lastEnd := 0

	for _, r := range ranges {
		if r.start > lastEnd {
			numIDs += r.end - r.start + 1
		} else if r.start == lastEnd {
			numIDs += r.end - r.start
		} else if r.end <= lastEnd {
			// do nothing
		} else {
			numIDs += r.end - lastEnd
		}
		lastStart = max(lastStart, r.start)
		lastEnd = max(lastEnd, r.end)
	}

	fmt.Println("Number of IDs: ", numIDs)
}
