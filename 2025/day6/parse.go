package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"os"
	"regexp"
)

func parseDataFromInput(lines []string) ([][]string, []string) {
	numberGroupPattern := regexp.MustCompile("([0-9]+)")
	operatorGroupPattern := regexp.MustCompile(`([+*-\/])`)

	var data [][]string
	var operators []string

	for _, line := range lines {
		matches := numberGroupPattern.FindAllString(line, -1)

		if len(matches) > 0 {
			data = append(data, matches)
		}
	}

	data = utils.SwapColumnsAndRows(data)

	for _, line := range lines {
		matches := operatorGroupPattern.FindAllString(line, -1)

		if len(matches) > 0 {
			for _, match := range matches {
				operators = append(operators, match)
			}
		}
	}

	return data, operators
}

func part1(lines []string) int64 {
	var result int64 = 0

	data, operators := parseDataFromInput(lines)

	for dataIndex, values := range data {
		operator := operators[dataIndex]

		switch operator {
		case "+":
			result += utils.SumArray(utils.ParseIntArray(values))
		case "*":
			result += utils.MultiplyArray(utils.ParseIntArray(values))
		default:
			fmt.Println("unknown operator")
		}
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

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}
