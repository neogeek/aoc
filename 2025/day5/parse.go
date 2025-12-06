package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var rangePattern = regexp.MustCompile(`^[0-9]+\-[0-9]+$`)
var idPattern = regexp.MustCompile(`^[0-9]+$`)

func parseDataFromInput(lines []string) ([]utils.Range, []uint64) {
	var ingredientRanges []utils.Range
	var ids []uint64

	for _, line := range lines {
		if rangePattern.MatchString(line) {
			parts := strings.Split(line, "-")
			start, _ := strconv.ParseUint(parts[0], 10, 64)
			end, _ := strconv.ParseUint(parts[1], 10, 64)

			ingredientRanges = append(ingredientRanges, utils.Range{Start: start, End: end})
		} else if idPattern.MatchString(line) {
			id, _ := strconv.ParseUint(line, 10, 64)

			ids = append(ids, id)
		}
	}

	sort.Slice(ingredientRanges, func(i, j int) bool {
		return ingredientRanges[i].Start < ingredientRanges[j].Start
	})

	return ingredientRanges, ids
}

func part1(lines []string) int {
	var result = 0

	ingredientRanges, ids := parseDataFromInput(lines)

	for _, id := range ids {
		var timesFoundInRange = 0
		for _, ingredientRange := range ingredientRanges {
			if id >= ingredientRange.Start && id <= ingredientRange.End {
				timesFoundInRange += 1
			}
		}
		if timesFoundInRange > 0 {
			result += 1
		}
	}

	return result
}

func part2(lines []string) uint64 {
	var result uint64 = 0

	ingredientRanges, _ := parseDataFromInput(lines)

	var combinedIngredientRanges []utils.Range

	combinedIngredientRanges = append(combinedIngredientRanges, utils.Range{
		Start: ingredientRanges[0].Start,
		End:   ingredientRanges[0].End,
	})

	for i := 1; i < len(ingredientRanges); i += 1 {
		length := len(combinedIngredientRanges) - 1

		if ingredientRanges[i].Start <= combinedIngredientRanges[length].End {
			combinedIngredientRanges[length].End = max(ingredientRanges[i].End, combinedIngredientRanges[length].End)
		} else {
			combinedIngredientRanges = append(combinedIngredientRanges, utils.Range{
				Start: ingredientRanges[i].Start,
				End:   ingredientRanges[i].End,
			})
		}
	}

	for _, combinedIngredientRange := range combinedIngredientRanges {
		result += (combinedIngredientRange.End - combinedIngredientRange.Start) + 1
	}

	fmt.Println(ingredientRanges)
	fmt.Println(combinedIngredientRanges)

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
