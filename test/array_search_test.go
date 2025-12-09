package dsa_brushup_test

import (
	"dsa-brushup"
	"testing"
)

func TestArraySearch(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		key      int
		expected bool
	}{
		{
			name:     "element found at beginning",
			input:    []int{1, 2, 3, 4, 5},
			key:      1,
			expected: true,
		},
		{
			name:     "element found at end",
			input:    []int{1, 2, 3, 4, 5},
			key:      5,
			expected: true,
		},
		{
			name:     "element found in middle",
			input:    []int{1, 2, 3, 4, 5},
			key:      3,
			expected: true,
		},
		{
			name:     "element not found",
			input:    []int{1, 2, 3, 4, 5},
			key:      10,
			expected: false,
		},
		{
			name:     "single element found",
			input:    []int{42},
			key:      42,
			expected: true,
		},
		{
			name:     "single element not found",
			input:    []int{42},
			key:      10,
			expected: false,
		},
		{
			name:     "empty array",
			input:    []int{},
			key:      5,
			expected: false,
		},
		{
			name:     "search for zero",
			input:    []int{1, 0, 2, 3},
			key:      0,
			expected: true,
		},
		{
			name:     "negative number found",
			input:    []int{-5, -3, -1, 0, 2},
			key:      -3,
			expected: true,
		},
		{
			name:     "negative number not found",
			input:    []int{-5, -3, -1, 0, 2},
			key:      -10,
			expected: false,
		},
		{
			name:     "duplicate elements",
			input:    []int{1, 2, 2, 3, 4},
			key:      2,
			expected: true,
		},
		{
			name:     "large array element found",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			key:      7,
			expected: true,
		},
		{
			name:     "large array element not found",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			key:      15,
			expected: false,
		},
		{
			name:     "all same elements found",
			input:    []int{5, 5, 5, 5},
			key:      5,
			expected: true,
		},
		{
			name:     "all same elements not found",
			input:    []int{5, 5, 5, 5},
			key:      3,
			expected: false,
		},
		{
			name:     "two elements first",
			input:    []int{1, 2},
			key:      1,
			expected: true,
		},
		{
			name:     "two elements second",
			input:    []int{1, 2},
			key:      2,
			expected: true,
		},
		{
			name:     "two elements not found",
			input:    []int{1, 2},
			key:      3,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dsa_brushup.Array_serach(tt.input, tt.key)
			if result != tt.expected {
				t.Errorf("Array_serach(%v, %d) = %v; want %v", tt.input, tt.key, result, tt.expected)
			}
		})
	}
}
