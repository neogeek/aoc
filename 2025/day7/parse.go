package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"os"
	"strings"
)

func part1(lines []string) int {
	var splits = 0

	data := utils.ExtractRowsAndColumns(lines, "")

	for y := 0; y < len(data); y += 1 {
		for x := 0; x < len(data[y]); x += 1 {
			switch data[y][x] {
			case "S":
				if data[y+1][x] == "." {
					data[y+1][x] = "|"
				}
			case "^":
				if x > 0 && data[y][x-1] == "." {
					data[y][x-1] = "|"

					if len(data) > y+1 && data[y+1][x-1] == "." {
						data[y+1][x-1] = "|"
					}
				}
				if len(data[y]) > x+1 && data[y][x+1] == "." {
					data[y][x+1] = "|"

					if len(data) > y+1 && data[y+1][x+1] == "." {
						data[y+1][x+1] = "|"
					}
				}

				if data[y-1][x] == "|" {
					splits += 1
				}
			case "|":
				if len(data) > y+1 && data[y+1][x] == "." {
					data[y+1][x] = "|"
				}
			}
		}
	}

	for _, row := range data {
		fmt.Println(strings.Join(row, ""))
	}

	return splits
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
