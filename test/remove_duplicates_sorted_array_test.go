package dsa_brushup_test

import (
	"dsa-brushup"
	"reflect"
	"testing"
)

func TestRemoveDuplicatesSortedArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "array with duplicates",
			input:    []int{1, 1, 2, 2, 3, 4, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "no duplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "all same elements",
			input:    []int{5, 5, 5, 5, 5},
			expected: []int{5},
		},
		{
			name:     "two elements same",
			input:    []int{1, 1},
			expected: []int{1},
		},
		{
			name:     "two elements different",
			input:    []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "multiple consecutive duplicates",
			input:    []int{1, 1, 1, 2, 2, 2, 3, 3, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "negative numbers with duplicates",
			input:    []int{-5, -5, -3, -1, -1, 0, 0, 2},
			expected: []int{-5, -3, -1, 0, 2},
		},
		{
			name:     "duplicates at beginning",
			input:    []int{1, 1, 1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "duplicates at end",
			input:    []int{1, 2, 3, 4, 4, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "alternating duplicates",
			input:    []int{1, 1, 2, 2, 3, 3, 4, 4},
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dsa_brushup.Remove_duplicates_sorted_array(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Remove_duplicates_sorted_array(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
