package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"os"
	"strings"
)

func part1(lines []string) int {
	var result = 0

	var positions []utils.Vector3

	for _, line := range lines {
		values := utils.ParseFloatArray(strings.Split(line, ","))

		positions = append(positions, utils.Vector3{X: values[0], Y: values[1], Z: values[2]})
	}

	for _, position := range positions {
		fmt.Println(position)
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

	if strings.HasSuffix(path, "example.txt") {
		utils.Assert(part1 == 40, fmt.Sprintf("Part 1 = %v", part1))
		utils.Assert(part2 == -1, fmt.Sprintf("Part 2 = %v", part2))
	} else {
		fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
	}
}
