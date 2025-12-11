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
	"time"
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

func CalculateBoundingBox(a Vector2, b Vector2) BoundingBox {
	var minX = math.Min(a.X, b.X)
	var maxX = math.Max(a.X, b.X)

	var minY = math.Min(a.Y, b.Y)
	var maxY = math.Max(a.Y, b.Y)

	return BoundingBox{MinX: minX, MaxX: maxX, MinY: minY, MaxY: maxY}
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

const Epsilon = 1e-9

func IsPointOnSegment(point Vector2, leftVertex, rightVertex Vector2) bool {
	crossProduct := point.Subtract(leftVertex).Cross(rightVertex.Subtract(leftVertex))

	if math.Abs(crossProduct) > Epsilon {
		return false
	}

	xMin := math.Min(leftVertex.X, rightVertex.X)
	xMax := math.Max(leftVertex.X, rightVertex.X)
	yMin := math.Min(leftVertex.Y, rightVertex.Y)
	yMax := math.Max(leftVertex.Y, rightVertex.Y)

	isBoundedX := point.X >= xMin-Epsilon && point.X <= xMax+Epsilon
	isBoundedY := point.Y >= yMin-Epsilon && point.Y <= yMax+Epsilon

	return isBoundedX && isBoundedY

}

func IsPointInPolygon(point Vector2, vertices []Vector2) bool {
	verticesCount := len(vertices)

	if verticesCount < 3 {
		return false
	}

	windingNumber := 0.0

	for i := range verticesCount {
		leftVertex := vertices[i]
		rightVertex := vertices[(i+1)%verticesCount]

		if IsPointOnSegment(point, leftVertex, rightVertex) {
			return true
		}

		u := leftVertex.Subtract(point)
		v := rightVertex.Subtract(point)

		cross := u.Cross(v)
		dot := u.Dot(v)

		angle := math.Atan2(cross, dot)

		windingNumber += angle
	}

	if math.Abs(math.Abs(windingNumber)-2*math.Pi) < Epsilon {
		return true
	}

	if math.Abs(windingNumber) < Epsilon {
		return false
	}

	return math.Abs(windingNumber) > Epsilon
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
		return values[i] > values[j]
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

func ElapsedTimer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
