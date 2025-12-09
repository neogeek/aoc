package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"math"
	"os"
	"strings"
)

func parsePositionsFromInput(lines []string) []utils.Vector2 {
	var positions []utils.Vector2

	for _, line := range lines {
		values := utils.ParseFloatArray(strings.Split(line, ","))

		positions = append(positions, utils.Vector2{X: values[0], Y: values[1]})
	}

	return positions
}

func calculateRectangle(a utils.Vector2, b utils.Vector2) float64 {
	width := math.Abs(a.X-b.X) + 1
	height := math.Abs(a.Y-b.Y) + 1

	return width * height
}

func part1(lines []string) int {
	positions := parsePositionsFromInput(lines)

	var largestRectangle float64 = 0

	for _, a := range positions {
		for _, b := range positions {
			currentRectangle := calculateRectangle(a, b)

			if currentRectangle > largestRectangle {
				largestRectangle = currentRectangle
			}
		}
	}

	return int(largestRectangle)
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
