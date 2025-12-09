package dsa_brushup_test

import (
	"dsa-brushup"
	"testing"
)

func TestFindMissingNumberInArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "missing number at beginning",
			input:    []int{2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "missing number at end",
			input:    []int{1, 2, 3, 4},
			expected: 5,
		},
		{
			name:     "missing number in middle",
			input:    []int{1, 2, 4, 5},
			expected: 3,
		},
		{
			name:     "array with 1 to 10, missing 7",
			input:    []int{1, 2, 3, 4, 5, 6, 8, 9, 10},
			expected: 7,
		},
		{
			name:     "array with 1 to 6, missing 3",
			input:    []int{1, 2, 4, 5, 6},
			expected: 3,
		},
		{
			name:     "two elements missing first",
			input:    []int{2},
			expected: 1,
		},
		{
			name:     "two elements missing second",
			input:    []int{1},
			expected: 2,
		},
		{
			name:     "three elements missing middle",
			input:    []int{1, 3},
			expected: 2,
		},
		{
			name:     "large array missing 50",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 51},
			expected: 50,
		},
		{
			name:     "unsorted array missing 4",
			input:    []int{5, 3, 1, 2},
			expected: 4,
		},
		{
			name:     "unsorted array missing 2",
			input:    []int{3, 1, 4, 5},
			expected: 2,
		},
		{
			name:     "array 1 to 7 missing 5",
			input:    []int{1, 2, 3, 4, 6, 7},
			expected: 5,
		},
		{
			name:     "array 1 to 8 missing 1",
			input:    []int{2, 3, 4, 5, 6, 7, 8},
			expected: 1,
		},
		{
			name:     "array 1 to 8 missing 8",
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dsa_brushup.Find_missing_number_in_array(tt.input)
			if result != tt.expected {
				t.Errorf("Find_missing_number_in_array(%v) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
