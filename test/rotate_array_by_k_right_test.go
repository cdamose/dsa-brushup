package dsa_brushup_test

import (
	"dsa-brushup"
	"reflect"
	"testing"
)

func TestRotateArrayByKRight(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		k        int
		expected []int
	}{
		{
			name:     "rotate by 1",
			input:    []int{1, 2, 3, 4, 5},
			k:        1,
			expected: []int{5, 1, 2, 3, 4},
		},
		{
			name:     "rotate by 2",
			input:    []int{1, 2, 3, 4, 5},
			k:        2,
			expected: []int{4, 5, 1, 2, 3},
		},
		{
			name:     "rotate by 3",
			input:    []int{1, 2, 3, 4, 5},
			k:        3,
			expected: []int{3, 4, 5, 1, 2},
		},
		{
			name:     "rotate by array length",
			input:    []int{1, 2, 3, 4, 5},
			k:        5,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "rotate by 0",
			input:    []int{1, 2, 3, 4, 5},
			k:        0,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "rotate by more than length",
			input:    []int{1, 2, 3, 4, 5},
			k:        7,
			expected: []int{4, 5, 1, 2, 3},
		},
		{
			name:     "rotate by multiple of length",
			input:    []int{1, 2, 3, 4, 5},
			k:        10,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "single element",
			input:    []int{42},
			k:        1,
			expected: []int{42},
		},
		{
			name:     "two elements rotate by 1",
			input:    []int{1, 2},
			k:        1,
			expected: []int{2, 1},
		},
		{
			name:     "empty array",
			input:    []int{},
			k:        3,
			expected: []int{},
		},
		{
			name:     "negative numbers",
			input:    []int{-5, -3, -1, 0, 2},
			k:        2,
			expected: []int{0, 2, -5, -3, -1},
		},
		{
			name:     "large rotation",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			k:        4,
			expected: []int{7, 8, 9, 10, 1, 2, 3, 4, 5, 6},
		},
		{
			name:     "rotate by length minus 1",
			input:    []int{1, 2, 3, 4, 5},
			k:        4,
			expected: []int{2, 3, 4, 5, 1},
		},
		{
			name:     "three elements",
			input:    []int{1, 2, 3},
			k:        1,
			expected: []int{3, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dsa_brushup.Rotate_array_by_k_right(tt.input, tt.k)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Rotate_array_by_k_right(%v, %d) = %v; want %v", tt.input, tt.k, result, tt.expected)
			}
		})
	}
}
