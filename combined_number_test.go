package main

import (
	"testing"
)

func TestCombinedNumber(t *testing.T) {
	tests := []struct {
		input    []int
		expected string
	}{
		{[]int{50, 2, 1, 9}, "95021"},
		{[]int{5, 50, 56}, "56550"},
		{[]int{420, 42, 423}, "42423420"},
		{[]int{0, 0}, "0"},
	}

	for _, tt := range tests {
		actual := combined_number(tt.input)
		if actual != tt.expected {
			t.Errorf("combined_number(%v) = %v; expected %v", tt.input, actual, tt.expected)
		}
	}
}
