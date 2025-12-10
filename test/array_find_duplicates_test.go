package dsa_brushup_test

import (
	"dsa-brushup"
	"reflect"
	"sort"
	"testing"
)

func TestArrayDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "single duplicate",
			input:    []int{1, 2, 3, 2, 4},
			expected: []int{2},
		},
		{
			name:     "multiple duplicates",
			input:    []int{1, 2, 3, 2, 4, 3, 5},
			expected: []int{2, 3},
		},
		{
			name:     "no duplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{},
		},
		{
			name:     "all same elements",
			input:    []int{5, 5, 5, 5},
			expected: []int{5},
		},
		{
			name:     "two elements duplicate",
			input:    []int{1, 1},
			expected: []int{1},
		},
		{
			name:     "two elements no duplicate",
			input:    []int{1, 2},
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: []int{},
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "multiple occurrences of same number",
			input:    []int{1, 1, 1, 2, 2, 3},
			expected: []int{1, 2},
		},
		{
			name:     "duplicates at beginning",
			input:    []int{1, 1, 2, 3, 4},
			expected: []int{1},
		},
		{
			name:     "duplicates at end",
			input:    []int{1, 2, 3, 4, 4},
			expected: []int{4},
		},
		{
			name:     "alternating duplicates",
			input:    []int{1, 2, 1, 2, 3},
			expected: []int{1, 2},
		},
		{
			name:     "negative numbers with duplicates",
			input:    []int{-1, -2, -1, 0, 1, 1},
			expected: []int{-1, 1},
		},
		{
			name:     "large array with duplicates",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 5, 3, 1},
			expected: []int{1, 3, 5},
		},
		{
			name:     "consecutive duplicates",
			input:    []int{1, 1, 2, 2, 3, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "three of same number",
			input:    []int{1, 2, 1, 3, 1},
			expected: []int{1},
		},
		{
			name:     "zeros with duplicates",
			input:    []int{0, 1, 0, 2, 3},
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dsa_brushup.Array_duplicates(tt.input)
			
			// Sort both slices for comparison (order doesn't matter for duplicates)
			sort.Ints(result)
			sort.Ints(tt.expected)
			
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Array_duplicates(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
