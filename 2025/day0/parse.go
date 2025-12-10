package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"os"
)

func part1(lines []string) int64 {
	var result int64 = 0

	for _, line := range lines {
		fmt.Println(line)
	}

	return result
}

func part2(lines []string) int64 {
	var result int64 = 0

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
