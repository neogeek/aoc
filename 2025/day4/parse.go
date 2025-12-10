package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"os"
	"strings"
)

func countAdjacentTiles(grid [][]string, rowIndex int, colIndex int, pattern string) int {

	rowLength := len(grid)
	colLength := len(grid[0])

	var count = 0

	// Left (broken)
	if colIndex > 0 && grid[rowIndex][colIndex-1] == pattern {
		count += 1
	}

	// Right
	if colIndex < colLength-1 && grid[rowIndex][colIndex+1] == pattern {
		count += 1
	}

	// Top (broken)
	if rowIndex > 0 && grid[rowIndex-1][colIndex] == pattern {
		count += 1
	}

	// Bottom
	if rowIndex < rowLength-1 && grid[rowIndex+1][colIndex] == pattern {
		count += 1
	}

	// Top Left
	if rowIndex > 0 && colIndex > 0 && grid[rowIndex-1][colIndex-1] == pattern {
		count += 1
	}

	// Top Right
	if rowIndex > 0 && colIndex < colLength-1 && grid[rowIndex-1][colIndex+1] == pattern {
		count += 1
	}

	// Bottom Left
	if rowIndex < rowLength-1 && colIndex > 0 && grid[rowIndex+1][colIndex-1] == pattern {
		count += 1
	}

	// Bottom Right
	if rowIndex < rowLength-1 && colIndex < colLength-1 && grid[rowIndex+1][colIndex+1] == pattern {
		count += 1
	}

	return count
}

func findAllAdjacentTiles(grid [][]string, pattern string) []utils.Vector2 {
	var tiles []utils.Vector2

	for rowIndex, row := range grid {
		for colIndex, col := range row {
			if col == pattern {
				count := countAdjacentTiles(grid, rowIndex, colIndex, pattern)

				if count < 4 {
					tiles = append(tiles, utils.Vector2{X: float64(colIndex), Y: float64(rowIndex)})
				}
			}
		}
	}

	return tiles
}

func part1(lines []string) int {
	var result = 0

	var grid [][]string
	var displayGrid [][]string

	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	for _, line := range lines {
		displayGrid = append(displayGrid, strings.Split(line, ""))
	}

	tiles := findAllAdjacentTiles(grid, "@")

	result += len(tiles)

	for _, tile := range tiles {
		displayGrid[int(tile.Y)][int(tile.X)] = "X"
	}

	// for _, row := range displayGrid {
	// 	fmt.Println(row)
	// }

	return result
}

func part2(lines []string) int {
	var result = 0

	var grid [][]string

	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	for {
		tiles := findAllAdjacentTiles(grid, "@")

		count := len(tiles)

		result += count

		for _, tile := range tiles {
			grid[int(tile.Y)][int(tile.X)] = "."
		}

		if count == 0 {
			break
		}
	}

	for _, row := range grid {
		fmt.Println(row)
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
		utils.Assert(part1 == 13, fmt.Sprintf("Part 1 = %v", part1))
		utils.Assert(part2 == 43, fmt.Sprintf("Part 2 = %v", part2))
	} else {
		fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
	}
}
