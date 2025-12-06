package utils

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

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

func ParseIntArray(values []string) []int64 {
	results := make([]int64, 0, len(values))

	for _, s := range values {
		i, err := strconv.ParseInt(s, 10, 64)

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
		f, err := strconv.ParseFloat(s, 64)

		if err != nil {
			fmt.Printf("Error converting string '%s' to float: %v\n", s, err)
			continue
		}

		results = append(results, f)
	}

	return results
}

func SumArray(values []int64) int64 {
	var result int64 = 0

	for _, value := range values {
		result += value
	}

	return result
}

func MultiplyArray(values []int64) int64 {
	var result int64 = 1

	for _, value := range values {
		result *= value
	}

	return result
}
