package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func doesContainRepeatedDigits(value string, length float64) bool {
	if utils.HasDecimal(length) {
		return false
	}

	reWithGroups := regexp.MustCompile(fmt.Sprintf(`([0-9]{%d})`, int64(length)))

	matches := reWithGroups.FindAllString(value, -1)

	return len(matches) > 1 && utils.AllEqual(matches)
}

func doesContainComplexRepeatedDigits(value string) bool {
	var length = float64(len(value))
	var checks = length

	for checks > 0 {
		if doesContainRepeatedDigits(value, length/checks) {
			return true
		}

		checks -= 1
	}

	return false
}

func part1(lines []string) int {
	var total = 0

	for _, line := range lines {
		parts := strings.Split(line, "-")
		var start, _ = strconv.Atoi(parts[0])
		var end, _ = strconv.Atoi(parts[1])

		numbers := utils.MakeRange(start, end+1)

		for _, v := range numbers {
			length := strconv.Itoa(v)
			if doesContainRepeatedDigits(length, float64(len(length))/2) {
				total += v
			}
		}
	}

	return total
}

func part2(lines []string) int {
	var total = 0

	for _, line := range lines {
		parts := strings.Split(line, "-")
		var start, _ = strconv.Atoi(parts[0])
		var end, _ = strconv.Atoi(parts[1])

		numbers := utils.MakeRange(start, end+1)

		for _, v := range numbers {
			if doesContainComplexRepeatedDigits(strconv.Itoa(v)) {
				total += v
			}
		}
	}

	return total
}

func main() {
	path := os.Args[1]

	lines, err := utils.LoadInput(path, ",")

	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	part1 := part1(lines)
	part2 := part2(lines)

	if strings.HasSuffix(path, "example.txt") {
		utils.Assert(part1 == 1227775554, fmt.Sprintf("Part 1 = %v", part1))
		utils.Assert(part2 == 4174379265, fmt.Sprintf("Part 2 = %v", part2))
	} else {
		fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
	}
}
