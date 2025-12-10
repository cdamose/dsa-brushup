package dsa_brushup_test

import (
	"dsa-brushup"
	"reflect"
	"testing"
)

func TestArraySortColor(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "mixed colors",
			input:    []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "already sorted",
			input:    []int{0, 0, 1, 1, 2, 2},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "reverse sorted",
			input:    []int{2, 2, 1, 1, 0, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "all zeros",
			input:    []int{0, 0, 0, 0},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "all ones",
			input:    []int{1, 1, 1, 1},
			expected: []int{1, 1, 1, 1},
		},
		{
			name:     "all twos",
			input:    []int{2, 2, 2, 2},
			expected: []int{2, 2, 2, 2},
		},
		{
			name:     "single element zero",
			input:    []int{0},
			expected: []int{0},
		},
		{
			name:     "single element one",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "single element two",
			input:    []int{2},
			expected: []int{2},
		},
		{
			name:     "two elements",
			input:    []int{2, 0},
			expected: []int{0, 2},
		},
		{
			name:     "only zeros and ones",
			input:    []int{1, 0, 1, 0, 1, 0},
			expected: []int{0, 0, 0, 1, 1, 1},
		},
		{
			name:     "only zeros and twos",
			input:    []int{2, 0, 2, 0, 2, 0},
			expected: []int{0, 0, 0, 2, 2, 2},
		},
		{
			name:     "only ones and twos",
			input:    []int{2, 1, 2, 1, 2, 1},
			expected: []int{1, 1, 1, 2, 2, 2},
		},
		{
			name:     "alternating pattern",
			input:    []int{0, 1, 2, 0, 1, 2},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "zeros at end",
			input:    []int{2, 1, 1, 2, 0, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "twos at beginning",
			input:    []int{2, 2, 0, 1, 0, 1},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "large array",
			input:    []int{2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1},
			expected: []int{0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2},
		},
		{
			name:     "mostly zeros",
			input:    []int{0, 0, 0, 1, 0, 2, 0},
			expected: []int{0, 0, 0, 0, 0, 1, 2},
		},
		{
			name:     "mostly twos",
			input:    []int{2, 2, 2, 1, 2, 0, 2},
			expected: []int{0, 1, 2, 2, 2, 2, 2},
		},
		{
			name:     "three elements",
			input:    []int{2, 0, 1},
			expected: []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dsa_brushup.ArraySortColorCountLogic(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ArraySortColor(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
