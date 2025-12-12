package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"math"
	"os"
	"sort"
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

func calculateArea(a utils.Vector2, b utils.Vector2) float64 {
	width := math.Abs(a.X-b.X) + 1
	height := math.Abs(a.Y-b.Y) + 1

	return width * height
}

func generateDebugGrid(positions []utils.Vector2) [][]string {
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

	grid := utils.MakeGrid(int(largestY)+2, int(largestX)+3, ".")

	for _, position := range positions {
		grid[int(position.Y)][int(position.X)] = "#"
	}

	for y := 0; y < len(grid); y += 1 {
		for x := 0; x < len(grid[y]); x += 1 {
			if utils.IsPointInPolygon(utils.Vector2{X: float64(x), Y: float64(y)}, positions) {
				if grid[y][x] != "#" {
					grid[y][x] = "X"
				}
			}
		}
	}

	return grid
}

func part1(lines []string) int64 {
	positions := parsePositionsFromInput(lines)

	var largestRectangle float64 = 0

	for _, a := range positions {
		for _, b := range positions {
			currentRectangle := calculateArea(a, b)

			if currentRectangle > largestRectangle {
				largestRectangle = currentRectangle
			}
		}
	}

	return int64(largestRectangle)
}

type Tile struct {
	Area float64
	Box  utils.BoundingBox
}

func IsTileInPolygon(boundingBox utils.BoundingBox, vertices []utils.Vector2) bool {
	p1 := utils.Vector2{X: boundingBox.MinX, Y: boundingBox.MinY}
	p2 := utils.Vector2{X: boundingBox.MaxX, Y: boundingBox.MinY}
	p3 := utils.Vector2{X: boundingBox.MinX, Y: boundingBox.MaxY}
	p4 := utils.Vector2{X: boundingBox.MaxX, Y: boundingBox.MaxY}

	if utils.IsPointInPolygon(p1, vertices) &&
		utils.IsPointInPolygon(p2, vertices) &&
		utils.IsPointInPolygon(p3, vertices) &&
		utils.IsPointInPolygon(p4, vertices) {
		return true
	}

	return false
}

func part2(lines []string) int64 {
	positions := parsePositionsFromInput(lines)

	var tiles []Tile

	for _, a := range positions {
		for _, b := range positions {
			if a != b && a.X < b.X && a.Y < b.Y {
				var boundingBox = utils.CalculateBoundingBox(a, b)
				var tile = Tile{Area: calculateArea(
					utils.Vector2{X: boundingBox.MinX, Y: boundingBox.MinY},
					utils.Vector2{X: boundingBox.MaxX, Y: boundingBox.MaxY},
				), Box: boundingBox}
				tiles = append(tiles, tile)
			}
		}
	}

	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i].Area > tiles[j].Area
	})

	for _, a := range tiles {
		var overlaps = true
		for _, b := range tiles {
			if a != b && !a.Box.Overlaps(b.Box) {
				overlaps = false
			}
		}

		fmt.Println(overlaps)
		fmt.Println(a)
	}

	return 0
}

func main() {
	defer utils.ElapsedTimer("Day 9")()

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
