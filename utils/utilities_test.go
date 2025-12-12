package utils

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestAllEqual(t *testing.T) {
	{
		input := []int{1, 2, 3, 4, 5}
		result := AllEqual(input)
		if result == true {
			t.Errorf(`AllEqual(%v) failed to return the correct result: %v`, input, result)
		}
	}
	{
		input := []int{1, 1, 1, 1, 1}
		result := AllEqual(input)
		if result == false {
			t.Errorf(`AllEqual(%v) failed to return the correct result: %v`, input, result)
		}
	}
}

func TestChunk(t *testing.T) {
	input := "12345"
	expected := []string{"12", "34", "5"}
	result := Chunk(input, 2)

	if slices.Equal(result, expected) == false {
		t.Errorf(`Chunk(%v) failed to return the correct result: %v`, input, result)
	}
}

func TestChunkWithVariableLength(t *testing.T) {
	input := "12345"
	expected := []string{"1", "234", "5 "}
	result := ChunkWithVariableLength(input, []int64{1, 3, 2}, 0)

	if slices.Equal(result, expected) == false {
		t.Errorf(`ChunkWithVariableLength(%v) failed to return the correct result: %v`, input, result)
	}
}

func TestDistanceBetweenVector2(t *testing.T) {
	{
		a := Vector2{0, 0}
		b := Vector2{1, 0}
		var expected float64 = 1
		result := DistanceBetweenVector2(a, b)

		if fmt.Sprintf("%.2f", result) != fmt.Sprintf("%.2f", expected) {
			t.Errorf(`DistanceBetweenVector2(%v, %v) failed to return the correct result: %v`, a, b, result)
		}
	}
	{
		a := Vector2{0, 0}
		b := Vector2{1, 1}
		var expected float64 = 1.41
		result := DistanceBetweenVector2(a, b)

		if fmt.Sprintf("%.2f", result) != fmt.Sprintf("%.2f", expected) {
			t.Errorf(`DistanceBetweenVector2(%v, %v) failed to return the correct result: %v`, a, b, result)
		}
	}
}

func TestDistanceBetweenVector3(t *testing.T) {
	{
		a := Vector3{0, 0, 0}
		b := Vector3{1, 0, 0}
		var expected float64 = 1
		result := DistanceBetweenVector3(a, b)

		if fmt.Sprintf("%.2f", result) != fmt.Sprintf("%.2f", expected) {
			t.Errorf(`DistanceBetweenVector3(%v, %v) failed to return the correct result: %v`, a, b, result)
		}
	}
	{
		a := Vector3{0, 0, 0}
		b := Vector3{1, 1, 1}
		var expected float64 = 1.73
		result := DistanceBetweenVector3(a, b)

		if fmt.Sprintf("%.2f", result) != fmt.Sprintf("%.2f", expected) {
			t.Errorf(`DistanceBetweenVector3(%v, %v) failed to return the correct result: %v`, a, b, result)
		}
	}
}

func TestExtractRowsAndColumns(t *testing.T) {
	input := []string{"1 23", "4 5 6"}
	expected := [][]string{{"1", "23"}, {"4", "5", "6"}}
	result := ExtractRowsAndColumns(input, `\s+`)

	if reflect.DeepEqual(result, expected) == false {
		t.Errorf(`ExtractRowsAndColumns(%v) failed to return the correct result: %v`, input, result)
	}
}

func TestExtractRowsAndColumnsWithWhitespace(t *testing.T) {
	input := []string{"123 456  78", "123 456 789"}
	expected := [][]string{{"123", "456", " 78"}, {"123", "456", "789"}}
	result := ExtractRowsAndColumnsWithWhitespace(input)

	if reflect.DeepEqual(result, expected) == false {
		t.Errorf(`ExtractRowsAndColumnsWithWhitespace(%v) failed to return the correct result: %v`, input, result)
	}
}

func TestHasDecimal(t *testing.T) {
	{
		var input float64 = 1
		result := HasDecimal(input)
		if HasDecimal(1) == true {
			t.Errorf(`HasDecimal(%v) failed to return the correct result: %v`, input, result)
		}
	}
	{
		var input float64 = 1.5
		result := HasDecimal(input)
		if result == false {
			t.Errorf(`HasDecimal(%v) failed to return the correct result: %v`, input, result)
		}
	}
}

func TestLoadInput(t *testing.T) {
	input := "../mocks/example.txt"
	lines, err := LoadInput(input, "\n")
	if err != nil || len(lines) == 0 || lines[0] != "line1" {
		t.Errorf(`LoadInput(%v) failed to return the correct result: %v`, input, err)
	}
}

func TestMakeRange(t *testing.T) {
	start := 10
	end := 20
	result := MakeRange(start, end)
	if len(result) != 10 {
		t.Errorf(`MakeRange(%v, %v) failed to return the correct result: %v`, start, end, result)
	}
}

func TestParseIntArray(t *testing.T) {
	input := []string{"1", "2", "3", "4", "5"}
	expected := []int64{1, 2, 3, 4, 5}
	result := ParseIntArray(input)
	if slices.Equal(result, expected) == false {
		t.Errorf(`ParseIntArray(%v) failed to return the correct result: %v`, input, result)
	}
}

func TestParseFloatArray(t *testing.T) {
	input := []string{"1", "1.25", "1.5", "1.75", "2"}
	expected := []float64{1, 1.25, 1.5, 1.75, 2}
	result := ParseFloatArray(input)
	if slices.Equal(result, expected) == false {
		t.Errorf(`ParseFloatArray(%v) failed to return the correct result: %v`, input, result)
	}
}

func TestParseSumArray(t *testing.T) {
	input := []int64{1, 2, 3, 4, 5}
	result := SumArray(input)
	if result != 15 {
		t.Errorf(`SumArray(%v) failed to return the correct result: %v`, input, result)
	}
}

func TestSwapColumnsAndRows(t *testing.T) {
	input := [][]string{{"1", "2", "3"}, {"4", "5", "6"}}
	expected := [][]string{{"1", "4"}, {"2", "5", "3", "6"}}
	result := SwapColumnsAndRows(input)

	if reflect.DeepEqual(result, expected) {
		t.Errorf(`SwapColumnsAndRows(%v) failed to return the correct result: %v`, input, result)
	}
}

func TestParseMultiplyArray(t *testing.T) {
	input := []int64{1, 2, 3, 4, 5}
	result := MultiplyArray(input)
	if result != 120 {
		t.Errorf(`MultiplyArray(%v) failed to return the correct result: %v`, input, result)
	}
}

func TestRemoveFromSlice(t *testing.T) {
	input := []int64{1, 2, 3, 4, 5}
	expected := []int64{1, 2, 4, 5}
	result := RemoveFromSlice(input, 3)

	if slices.Equal(result, expected) == false {
		t.Errorf(`RemoveFromSlice(%v) failed to return the correct result: %v`, input, result)
	}
}

func TestUniqueSlice(t *testing.T) {
	{
		input := []int64{1, 1, 2, 3, 4, 4, 5}
		expected := []int64{1, 2, 3, 4, 5}
		result := UniqueSlice(input)

		if slices.Equal(result, expected) == false {
			t.Errorf(`UniqueSlice(%v) failed to return the correct result: %v`, input, result)
		}
	}
	{
		input := []Vector3{Vector3{1, 2, 3}, Vector3{1, 2, 3}, Vector3{5, 5, 5}}
		expected := []Vector3{Vector3{1, 2, 3}, Vector3{5, 5, 5}}
		result := UniqueSlice(input)

		if slices.Equal(result, expected) == false {
			t.Errorf(`UniqueSlice(%v) failed to return the correct result: %v`, input, result)
		}
	}
}
