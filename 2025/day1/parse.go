package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func rotate(direction string, position int, amount int) (int, int) {
	var newPosition = position
	var remainingAmount = amount
	var rotations = 0

	if direction == "L" {
		remainingAmount = -remainingAmount
	}

	if remainingAmount > 0 {
		for remainingAmount > 0 {
			remainingAmount -= 1
			newPosition -= 1

			if newPosition == -1 {
				newPosition = 99
			}

			if newPosition == 0 {
				rotations += 1
			}
		}
	} else if remainingAmount < 0 {
		for remainingAmount < 0 {
			remainingAmount += 1
			newPosition += 1

			if newPosition == 100 {
				newPosition = 0
			}

			if newPosition == 0 {
				rotations += 1
			}
		}
	}

	return newPosition, rotations
}

var reWithGroups = regexp.MustCompile(`([RL])([0-9]+)`)

func part1(lines []string) int {
	var position = 50
	var totalRotations = 0

	for _, line := range lines {
		submatches := reWithGroups.FindStringSubmatch(line)

		if len(submatches) != 3 {
			continue
		}

		amount, _ := strconv.Atoi(submatches[2])

		newPosition, _ := rotate(submatches[1], position, amount)

		position = newPosition

		if position == 0 {
			totalRotations += 1
		}
	}

	return totalRotations
}

func part2(lines []string) int {
	var position = 50
	var totalRotations = 0

	for _, line := range lines {
		submatches := reWithGroups.FindStringSubmatch(line)

		if len(submatches) != 3 {
			continue
		}

		amount, _ := strconv.Atoi(submatches[2])

		newPosition, rotations := rotate(submatches[1], position, amount)

		position = newPosition

		totalRotations += rotations
	}

	return totalRotations
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
