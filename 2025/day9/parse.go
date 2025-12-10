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

	positions := parsePositionsFromInput(lines)

	var largestX float64 = 0
	var largestY float64 = 0

	for _, position := range positions {
		if position.X > largestX {
			largestX = position.X
		}
		if position.Y > largestY {
			largestY = position.Y
		}
	}

	var grid [][]string

	for rowIndex := range int(largestY + 3) {
		var row = make([]string, int(largestX+3))

		for colIndex, _ := range row {
			row[colIndex] = "."

			for _, position := range positions {
				if rowIndex == int(position.Y) && colIndex == int(position.X) {
					row[colIndex] = "#"
				}
			}
		}

		grid = append(grid, row)
	}

	for _, line := range grid {
		fmt.Println(strings.Join(line, ""))
	}

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
		utils.Assert(part1 == 50, fmt.Sprintf("Part 1 = %v", part1))
		utils.Assert(part2 == 24, fmt.Sprintf("Part 2 = %v", part2))
	} else {
		fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
	}
}
