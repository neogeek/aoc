package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var rangePattern = regexp.MustCompile(`^[0-9]+\-[0-9]+$`)
var idPattern = regexp.MustCompile(`^[0-9]+$`)

func part1(lines []string) int {
	var result = 0

	var ingredientRanges []utils.Range
	var ids []uint64

	for _, line := range lines {
		if rangePattern.MatchString(line) {
			parts := strings.Split(line, "-")
			start, _ := strconv.ParseUint(parts[0], 10, 64)
			end, _ := strconv.ParseUint(parts[1], 10, 64)

			ingredientRanges = append(ingredientRanges, utils.Range{Start: start, End: end})
		} else if idPattern.MatchString(line) {
			id, _ := strconv.ParseUint(line, 10, 64)

			ids = append(ids, id)
		}
	}

	for _, id := range ids {
		var timesFoundInRange = 0
		for _, ingredientRange := range ingredientRanges {
			if id >= ingredientRange.Start && id <= ingredientRange.End {
				timesFoundInRange += 1
			}
		}
		if timesFoundInRange > 0 {
			result += 1
		}
	}

	return result
}

func part2(lines []string) int {
	var result = 0

	return result
}

func main() {
	path := os.Args[1]

	lines, err := utils.LoadInput(path, "\n")

	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	part1 := part1(lines)
	part2 := part2(lines)

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}
