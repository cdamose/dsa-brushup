package dsa_brushup_test

import (
	"dsa-brushup"
	"testing"
)

func TestMaximumSubArrayKadaneAlgorithm(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "classic example",
			input:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6, // [4, -1, 2, 1]
		},
		{
			name:     "all positive",
			input:    []int{1, 2, 3, 4, 5},
			expected: 15, // entire array
		},
		{
			name:     "all negative",
			input:    []int{-5, -2, -8, -1, -4},
			expected: -1, // single element [-1]
		},
		{
			name:     "single element positive",
			input:    []int{5},
			expected: 5,
		},
		{
			name:     "single element negative",
			input:    []int{-3},
			expected: -3,
		},
		{
			name:     "two elements positive",
			input:    []int{3, 5},
			expected: 8,
		},
		{
			name:     "two elements negative",
			input:    []int{-3, -5},
			expected: -3,
		},
		{
			name:     "two elements mixed",
			input:    []int{-3, 5},
			expected: 5,
		},
		{
			name:     "alternating positive negative",
			input:    []int{1, -2, 3, -4, 5},
			expected: 5, // [5]
		},
		{
			name:     "max at beginning",
			input:    []int{5, -3, -2, -1},
			expected: 5,
		},
		{
			name:     "max at end",
			input:    []int{-3, -2, -1, 5},
			expected: 5,
		},
		{
			name:     "max in middle",
			input:    []int{-2, 3, 4, -1},
			expected: 7, // [3, 4]
		},
		{
			name:     "entire array is max",
			input:    []int{1, 2, 3, 4},
			expected: 10,
		},
		{
			name:     "zeros in array",
			input:    []int{-1, 0, -2, 0, 3},
			expected: 3,
		},
		{
			name:     "all zeros",
			input:    []int{0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "large positive in middle",
			input:    []int{-2, -3, 10, -1, -2},
			expected: 10,
		},
		{
			name:     "multiple subarrays same sum",
			input:    []int{1, 2, -3, 4, 5, -9, 6},
			expected: 9, // [4, 5]
		},
		{
			name:     "negative then positive",
			input:    []int{-5, -4, -3, 10},
			expected: 10,
		},
		{
			name:     "positive then negative",
			input:    []int{10, -5, -4, -3},
			expected: 10,
		},
		{
			name:     "complex case",
			input:    []int{5, -3, 5, -2, 3, -1, 4},
			expected: 11, // [5, -3, 5, -2, 3, -1, 4]
		},
		{
			name:     "kadane edge case - reset needed",
			input:    []int{-2, -3, 4, -1, -2, 1, 5, -3},
			expected: 7, // [4, -1, -2, 1, 5]
		},
		{
			name:     "large array",
			input:    []int{1, -1, 2, -1, 3, -1, 4, -1, 5},
			expected: 11, // entire array: 1-1+2-1+3-1+4-1+5 = 11
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dsa_brushup.Maximum_sub_array_kadane_algorithm(tt.input)
			if result != tt.expected {
				t.Errorf("Maximum_sub_array_kadane_algorithm(%v) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
