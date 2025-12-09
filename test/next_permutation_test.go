package dsa_brushup_test

import (
	"dsa-brushup"
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "basic case",
			input:    []int{1, 2, 3},
			expected: []int{1, 3, 2},
		},
		{
			name:     "middle permutation",
			input:    []int{1, 3, 2},
			expected: []int{2, 1, 3},
		},
		{
			name:     "last permutation wraps to first",
			input:    []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
		{
			name:     "two elements ascending",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "two elements descending",
			input:    []int{2, 1},
			expected: []int{1, 2},
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "with duplicates",
			input:    []int{1, 5, 1},
			expected: []int{5, 1, 1},
		},
		{
			name:     "all same elements",
			input:    []int{1, 1, 1},
			expected: []int{1, 1, 1},
		},
		{
			name:     "larger array",
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 2, 4, 3},
		},
		{
			name:     "larger array near end",
			input:    []int{1, 4, 3, 2},
			expected: []int{2, 1, 3, 4},
		},
		{
			name:     "complex case",
			input:    []int{1, 3, 5, 4, 2},
			expected: []int{1, 4, 2, 3, 5},
		},
		{
			name:     "ascending order",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 5, 4},
		},
		{
			name:     "descending order",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "partial descending",
			input:    []int{2, 3, 1},
			expected: []int{3, 1, 2},
		},
		{
			name:     "with zeros",
			input:    []int{0, 1, 2},
			expected: []int{0, 2, 1},
		},
		{
			name:     "larger numbers",
			input:    []int{5, 6, 7},
			expected: []int{5, 7, 6},
		},
		{
			name:     "four elements last permutation",
			input:    []int{4, 3, 2, 1},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "swap at end",
			input:    []int{1, 2, 4, 3},
			expected: []int{1, 3, 2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dsa_brushup.Next_permutation(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Next_permutation(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
