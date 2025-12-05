package utils

import (
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
	{
		start := 10
		end := 20
		result := MakeRange(start, end)
		if len(result) != 10 {
			t.Errorf(`MakeRange(%v, %v) failed to return the correct result: %v`, start, end, result)
		}
	}
}

func TestParseIntArray(t *testing.T) {
	{
		input := []string{"1", "2", "3", "4", "5"}
		expected := []int64{1, 2, 3, 4, 5}
		result := ParseIntArray(input)
		if slices.Equal(result, expected) == false {
			t.Errorf(`ParseIntArray(%v) failed to return the correct result: %v`, input, result)
		}
	}
}

func TestParseFloatArray(t *testing.T) {
	{
		input := []string{"1", "1.25", "1.5", "1.75", "2"}
		expected := []float64{1, 1.25, 1.5, 1.75, 2}
		result := ParseFloatArray(input)
		if slices.Equal(result, expected) == false {
			t.Errorf(`ParseFloatArray(%v) failed to return the correct result: %v`, input, result)
		}
	}
}
