package dsa_brushup_test

import (
	"dsa-brushup"
	"testing"
)

func TestLargertElementInArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "positive numbers",
			input:    []int{1, 5, 3, 9, 2},
			expected: 9,
		},
		{
			name:     "negative numbers",
			input:    []int{-10, -5, -20, -1},
			expected: -1,
		},
		{
			name:     "mixed positive and negative",
			input:    []int{-5, 10, 3, -20, 7},
			expected: 10,
		},
		{
			name:     "single element",
			input:    []int{42},
			expected: 42,
		},
		{
			name:     "all same elements",
			input:    []int{5, 5, 5, 5},
			expected: 5,
		},
		{
			name:     "largest at beginning",
			input:    []int{100, 1, 2, 3},
			expected: 100,
		},
		{
			name:     "largest at end",
			input:    []int{1, 2, 3, 100},
			expected: 100,
		},
		{
			name:     "two elements",
			input:    []int{5, 10},
			expected: 10,
		},
		{
			name:     "with zero",
			input:    []int{0, -5, -10},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dsa_brushup.Largert_element_in_array(tt.input)
			if result != tt.expected {
				t.Errorf("Largert_element_in_array(%v) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
