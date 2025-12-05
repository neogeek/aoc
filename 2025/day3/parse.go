package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	var result int = 0

	for _, value := range lines {
		var batteries = utils.ParseIntArray(strings.Split(strings.TrimSpace(value), ""))

		var largestBattery int64 = 0

		for i := 0; i < len(batteries)-1; i += 1 {
			if batteries[i] > largestBattery {
				largestBattery = batteries[i]
			}
		}

		var indexOfLargestBattery = strings.Index(value, strconv.FormatInt(largestBattery, 10))

		var secondLargestBattery int64 = 0

		for i := indexOfLargestBattery + 1; i < len(batteries); i += 1 {
			if batteries[i] > secondLargestBattery {
				secondLargestBattery = batteries[i]
			}
		}

		output, _ := strconv.Atoi(fmt.Sprintf("%d%d", largestBattery, secondLargestBattery))

		result += output
	}

	return result
}

func part2(lines []string) int {
	var result int = 0

	// for _, value := range lines {
	// 	fmt.Println(value)
	// }

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
