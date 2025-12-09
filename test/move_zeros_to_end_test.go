package dsa_brushup_test

import (
	"dsa-brushup"
	"reflect"
	"testing"
)

func TestMoveZeroToEnd(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "zeros in middle",
			input:    []int{1, 0, 2, 0, 3, 4},
			expected: []int{1, 2, 3, 4, 0, 0},
		},
		{
			name:     "zeros at beginning",
			input:    []int{0, 0, 1, 2, 3},
			expected: []int{1, 2, 3, 0, 0},
		},
		{
			name:     "zeros at end",
			input:    []int{1, 2, 3, 0, 0},
			expected: []int{1, 2, 3, 0, 0},
		},
		{
			name:     "no zeros",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "all zeros",
			input:    []int{0, 0, 0, 0},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "single zero",
			input:    []int{0},
			expected: []int{0},
		},
		{
			name:     "single non-zero",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "alternating zeros",
			input:    []int{1, 0, 2, 0, 3, 0},
			expected: []int{1, 2, 3, 0, 0, 0},
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "two elements with zero first",
			input:    []int{0, 1},
			expected: []int{1, 0},
		},
		{
			name:     "two elements with zero last",
			input:    []int{1, 0},
			expected: []int{1, 0},
		},
		{
			name:     "negative numbers with zeros",
			input:    []int{-1, 0, -2, 0, 3},
			expected: []int{-1, -2, 3, 0, 0},
		},
		{
			name:     "multiple consecutive zeros",
			input:    []int{1, 0, 0, 0, 2, 3},
			expected: []int{1, 2, 3, 0, 0, 0},
		},
		{
			name:     "large array",
			input:    []int{1, 0, 2, 0, 3, 0, 4, 0, 5, 0},
			expected: []int{1, 2, 3, 4, 5, 0, 0, 0, 0, 0},
		},
		{
			name:     "preserve order of non-zeros",
			input:    []int{5, 0, 3, 0, 1, 0, 2},
			expected: []int{5, 3, 1, 2, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dsa_brushup.Move_zero_to_end(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Move_zero_to_end(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
