package utils

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func Assert(result bool, description string) {
	if !result {
		fmt.Printf("%s [FAILED]\n", description)
	} else {
		fmt.Printf("%s [OK]\n", description)
	}
}

func AllEqual[T comparable](slice []T) bool {
	if len(slice) <= 1 {
		return true
	}

	for i := 1; i < len(slice); i += 1 {
		if slice[i] != slice[0] {
			return false
		}
	}

	return true
}

func Chunk(value string, length int) []string {
	var result []string

	var chars = strings.Split(value, "")

	var chunk []string

	for i := 0; i < len(chars); i += 1 {
		chunk = append(chunk, chars[i])

		if len(chunk) == length {
			result = append(result, strings.Join(chunk, ""))

			chunk = []string{}
		}
	}

	if len(chunk) > 0 {
		result = append(result, strings.Join(chunk, ""))
	}

	return result
}

func ChunkWithVariableLength(value string, lengths []int64, padding int64) []string {
	var result []string

	var chars = strings.Split(value, "")

	var chunk []string

	var lengthIndex = 0

	for i := 0; i < len(chars); i += 1 {
		chunk = append(chunk, chars[i])

		if len(lengths) > lengthIndex && int64(len(chunk)) == lengths[lengthIndex] {
			result = append(result, strings.Join(chunk, ""))

			chunk = []string{}

			lengthIndex += 1

			i += int(padding)
		}
	}

	if len(chunk) > 0 {
		result = append(result, PadRight(strings.Join(chunk, ""), " ", int(lengths[lengthIndex])))
	}

	return result
}

func DistanceBetweenVector2(a Vector2, b Vector2) float64 {
	return math.Sqrt(math.Pow(b.X-a.X, 2) + math.Pow(b.Y-a.Y, 2))
}

func DistanceBetweenVector3(a Vector3, b Vector3) float64 {
	return math.Sqrt(math.Pow(b.X-a.X, 2) + math.Pow(b.Y-a.Y, 2) + math.Pow(b.Z-a.Z, 2))
}

func ExtractRowsAndColumns(lines []string, pattern string) [][]string {
	var result [][]string

	re := regexp.MustCompile(pattern)

	for _, line := range lines {
		result = append(result, re.Split(strings.TrimSpace(line), -1))
	}

	return result
}

func ExtractRowsAndColumnsWithWhitespace(lines []string) [][]string {
	var result [][]string

	data := ExtractRowsAndColumns(lines, `\s+`)

	data = SwapColumnsAndRows(data)

	maxColumnLengths := GetLengthOfColumns(data)

	data = SwapColumnsAndRows(data)

	for _, line := range lines {
		result = append(result, ChunkWithVariableLength(line, maxColumnLengths, 1))
	}

	return result
}

func GetLengthOfColumns(data [][]string) []int64 {
	var columnLengths []int64

	for _, row := range data {
		var maxColumnLength int64 = 0

		for _, col := range row {
			if int64(len(col)) > maxColumnLength {
				maxColumnLength = int64(len(col))
			}
		}

		columnLengths = append(columnLengths, maxColumnLength)
	}

	return columnLengths
}

func HasDecimal(num float64) bool {
	return math.Mod(num, 1.0) != 0.0
}

func LoadInput(filePath string, seperator string) ([]string, error) {
	content, err := os.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	return strings.Split(strings.TrimSpace(string(content)), seperator), nil
}

func MakeRange(start int, end int) []int {
	var length = end - start

	result := make([]int, length)

	for i := range result {
		result[i] = start + i
	}

	return result
}

func PadLeft(value string, char string, length int) string {

	if len(value) >= length {
		return value
	}

	return strings.Repeat(char, length-len(value)) + value
}

func PadRight(value string, char string, length int) string {

	if len(value) >= length {
		return value
	}

	return value + strings.Repeat(char, length-len(value))
}

func ParseIntArray(values []string) []int64 {
	results := make([]int64, 0, len(values))

	for _, s := range values {
		i, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)

		if err != nil {
			fmt.Printf("Error converting string '%s' to int: %v\n", s, err)
			continue
		}

		results = append(results, i)
	}

	return results
}

func ParseFloatArray(values []string) []float64 {
	results := make([]float64, 0, len(values))

	for _, s := range values {
		f, err := strconv.ParseFloat(strings.TrimSpace(s), 64)

		if err != nil {
			fmt.Printf("Error converting string '%s' to float: %v\n", s, err)
			continue
		}

		results = append(results, f)
	}

	return results
}

func Reverse[T Number](values []T) {
	sort.Slice(values, func(i, j int) bool {
		return i > j
	})
}

func SumArray[T Number](values []T) T {
	var result T = 0

	for _, value := range values {
		result += value
	}

	return result
}

func SwapColumnsAndRows(values [][]string) [][]string {
	var result [][]string

	if len(values) == 0 {
		return result
	}

	rowCount := len(values)
	colCount := len(values[0])

	result = make([][]string, colCount)

	for i := range colCount {
		result[i] = make([]string, rowCount)
	}

	for rowIndex, row := range values {
		for colIndex, col := range row {
			result[colIndex][rowIndex] = col
		}
	}

	return result
}

func MakeGrid(rowCount int, colCount int, fill string) [][]string {
	var grid [][]string

	for range rowCount {
		grid = append(grid, slices.Repeat([]string{"."}, colCount))
	}

	return grid
}

func MultiplyArray[T Number](values []T) T {
	var result T = 1

	for _, value := range values {
		result *= value
	}

	return result
}
