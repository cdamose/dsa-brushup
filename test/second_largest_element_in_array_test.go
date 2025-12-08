package dsa_brushup_test

import (
	"dsa-brushup"
	"testing"
)

func TestSecondLargestElementInArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "positive numbers",
			input:    []int{1, 5, 3, 9, 2},
			expected: 5,
		},
		{
			name:     "negative numbers",
			input:    []int{-10, -5, -20, -1},
			expected: -5,
		},
		{
			name:     "mixed positive and negative",
			input:    []int{-5, 10, 3, -20, 7},
			expected: 7,
		},
		{
			name:     "largest at beginning, second at end",
			input:    []int{100, 1, 2, 50},
			expected: 50,
		},
		{
			name:     "largest at end, second at beginning",
			input:    []int{50, 2, 3, 100},
			expected: 50,
		},
		{
			name:     "two elements",
			input:    []int{5, 10},
			expected: 5,
		},
		{
			name:     "with duplicates of largest",
			input:    []int{10, 10, 5, 3},
			expected: 10,
		},
		{
			name:     "all same elements",
			input:    []int{5, 5, 5, 5},
			expected: 5,
		},
		{
			name:     "three elements ascending",
			input:    []int{1, 2, 3},
			expected: 2,
		},
		{
			name:     "three elements descending",
			input:    []int{3, 2, 1},
			expected: 2,
		},
		{
			name:     "with zero",
			input:    []int{0, -5, 10, 5},
			expected: 5,
		},
		{
			name:     "second largest in middle",
			input:    []int{1, 8, 3, 10, 2},
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dsa_brushup.Second_largest_elemnt_in_array(tt.input)
			if result != tt.expected {
				t.Errorf("Second_largest_elemnt_in_array(%v) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
