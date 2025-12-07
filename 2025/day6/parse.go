package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"os"
	"strings"
)

func parseDataFromInput(lines []string) ([][]string, []string) {
	data := utils.ExtractRowsAndColumnsWithWhitespace(lines)
	operators := data[len(data)-1]
	data = data[:len(data)-1]

	data = utils.SwapColumnsAndRows(data)

	return data, operators
}

func part1(lines []string) int64 {
	var result int64 = 0

	data, operators := parseDataFromInput(lines)

	for dataIndex, values := range data {
		operator := strings.TrimSpace(operators[dataIndex])

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
