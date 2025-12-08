package dsa_brushup_test

import (
	"dsa-brushup"
	"testing"
)

func TestCheckArraySortedOrNot(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "sorted ascending",
			input:    []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "sorted with duplicates",
			input:    []int{1, 2, 2, 3, 4},
			expected: true,
		},
		{
			name:     "not sorted",
			input:    []int{1, 3, 2, 4, 5},
			expected: false,
		},
		{
			name:     "descending order",
			input:    []int{5, 4, 3, 2, 1},
			expected: false,
		},
		{
			name:     "single element",
			input:    []int{42},
			expected: true,
		},
		{
			name:     "two elements sorted",
			input:    []int{1, 2},
			expected: true,
		},
		{
			name:     "two elements not sorted",
			input:    []int{2, 1},
			expected: false,
		},
		{
			name:     "all same elements",
			input:    []int{5, 5, 5, 5},
			expected: true,
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: true,
		},
		{
			name:     "negative numbers sorted",
			input:    []int{-10, -5, -1, 0, 5},
			expected: true,
		},
		{
			name:     "negative numbers not sorted",
			input:    []int{-5, -10, -1, 0},
			expected: false,
		},
		{
			name:     "large sorted array",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: true,
		},
		{
			name:     "unsorted at end",
			input:    []int{1, 2, 3, 4, 3},
			expected: false,
		},
		{
			name:     "unsorted at beginning",
			input:    []int{2, 1, 3, 4, 5},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dsa_brushup.Check_array_sorted_or_not(tt.input)
			if result != tt.expected {
				t.Errorf("Check_array_sorted_or_not(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
