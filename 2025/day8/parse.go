package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

type JunctionBox struct {
	Positions []utils.Vector3
}

type Distance struct {
	Length    float64
	Positions []utils.Vector3
}

func (d Distance) Equal(other Distance) bool {
	if d.Length != other.Length {
		return false
	}

	if len(d.Positions) != len(other.Positions) {
		return false
	}

	for _, position := range other.Positions {
		if !slices.Contains(d.Positions, position) {
			return false
		}
	}

	return true
}

func (junctionBox JunctionBox) Contains(position utils.Vector3) bool {
	return slices.Contains(junctionBox.Positions, position)
}

func IsContainedInAnyJunctionBox(positions []utils.Vector3, junctionBoxes []JunctionBox) (bool, int) {
	for junctionBoxIndex, junctionBox := range junctionBoxes {
		if slices.ContainsFunc(positions, junctionBox.Contains) {
			return true, junctionBoxIndex
		}
	}

	return false, -1
}

func part1(lines []string) int {
	var result = 1

	var positions []utils.Vector3

	var finalDistances []Distance

	var junctionBoxes []JunctionBox

	for _, line := range lines {
		values := utils.ParseFloatArray(strings.Split(line, ","))

		positions = append(positions, utils.Vector3{X: values[0], Y: values[1], Z: values[2]})
	}

	for _, a := range positions {
		var distances []Distance

		for _, b := range positions {
			if a == b {
				continue
			}

			distance := Distance{Length: utils.DistanceBetweenVector3(a, b), Positions: []utils.Vector3{a, b}}

			var found bool = false

			for _, d := range finalDistances {
				if d.Equal(distance) {
					found = true
				}
			}

			if !found {
				distances = append(distances, distance)
			}
		}

		sort.Slice(distances, func(i, j int) bool {
			return distances[i].Length < distances[j].Length
		})

		if len(distances) <= 0 {
			break
		}

		var closestDistance = distances[0]

		finalDistances = append(finalDistances, closestDistance)
	}

	sort.Slice(finalDistances, func(i, j int) bool {
		return finalDistances[i].Length < finalDistances[j].Length
	})

	// for _, finalDistance := range finalDistances[0:11] {
	// 	fmt.Println(finalDistance)
	// }

	for _, finalDistance := range finalDistances[0:11] {
		foundInJunctionBox, junctionBoxIndex := IsContainedInAnyJunctionBox(finalDistance.Positions, junctionBoxes)

		if foundInJunctionBox {
			for _, position := range finalDistance.Positions {
				junctionBoxes[junctionBoxIndex].Positions = append(junctionBoxes[junctionBoxIndex].Positions, position)
			}
		} else {
			junctionBoxes = append(junctionBoxes, JunctionBox{Positions: finalDistance.Positions})
		}
	}

	for junctionBoxIndex, junctionBox := range junctionBoxes {
		junctionBoxes[junctionBoxIndex].Positions = utils.UniqueSlice(junctionBox.Positions)
	}

	sort.Slice(junctionBoxes, func(i, j int) bool {
		return len(junctionBoxes[i].Positions) > len(junctionBoxes[j].Positions)
	})

	// for _, junctionBox := range junctionBoxes {
	// 	fmt.Println(junctionBox)
	// }

	if len(junctionBoxes) > 3 {
		for _, junctionBox := range junctionBoxes[0:3] {
			result *= len(junctionBox.Positions)
		}
	}

	return result
}

func part2(lines []string) int {
	var result = 0

	return result
}

func main() {
	defer utils.ElapsedTimer("Day 8")()

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
