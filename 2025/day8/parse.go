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

	var junctionBoxes []JunctionBox

	for _, line := range lines {
		values := utils.ParseFloatArray(strings.Split(line, ","))

		positions = append(positions, utils.Vector3{X: values[0], Y: values[1], Z: values[2]})
	}

	for len(positions) > 0 {
		for _, a := range positions {
			var distances []Distance

			for _, b := range positions {
				if a == b {
					continue
				}

				distance := utils.DistanceBetweenVector3(a, b)

				if distance > 0 {
					distances = append(distances, Distance{Length: distance, Positions: []utils.Vector3{a, b}})
				}
			}

			for _, junctionBox := range junctionBoxes {
				for _, position := range junctionBox.Positions {
					distance := utils.DistanceBetweenVector3(a, position)

					distances = append(distances, Distance{Length: distance, Positions: []utils.Vector3{a, position}})
				}
			}

			sort.Slice(distances, func(i, j int) bool {
				return distances[i].Length < distances[j].Length
			})

			if len(distances) <= 0 {
				break
			}

			var closestDistance = distances[0]

			foundInJunctionBox, junctionBoxIndex := IsContainedInAnyJunctionBox(closestDistance.Positions, junctionBoxes)

			if foundInJunctionBox {
				for _, position := range closestDistance.Positions {
					junctionBoxes[junctionBoxIndex].Positions = append(junctionBoxes[junctionBoxIndex].Positions, position)
				}
			} else {
				junctionBoxes = append(junctionBoxes, JunctionBox{Positions: closestDistance.Positions})
			}

			for _, position := range closestDistance.Positions {
				positions = utils.RemoveFromSlice(positions, position)
			}
		}
	}

	for junctionBoxIndex, _ := range junctionBoxes {
		junctionBoxes[junctionBoxIndex].Positions = utils.UniqueSlice(junctionBoxes[junctionBoxIndex].Positions)
	}

	sort.Slice(junctionBoxes, func(i, j int) bool {
		return len(junctionBoxes[i].Positions) > len(junctionBoxes[j].Positions)
	})

	topThreeJunctionBoxes := junctionBoxes[0:3]

	for _, junctionBox := range topThreeJunctionBoxes {

		fmt.Println(len(junctionBox.Positions))

		result *= len(junctionBox.Positions)
	}

	for _, junctionBox := range topThreeJunctionBoxes {
		fmt.Println(junctionBox.Positions)
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
