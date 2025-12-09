package dsa_brushup_test

import (
	"dsa-brushup"
	"reflect"
	"testing"
)

func TestMatrixSetZero(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name: "single zero in middle",
			input: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			name: "multiple zeros",
			input: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			expected: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
		{
			name: "zero at top-left corner",
			input: [][]int{
				{0, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{0, 0, 0},
				{0, 1, 1},
				{0, 1, 1},
			},
		},
		{
			name: "zero at bottom-right corner",
			input: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 0},
			},
			expected: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 0},
			},
		},
		{
			name: "no zeros",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
		{
			name: "all zeros",
			input: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			expected: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "single row with zero",
			input: [][]int{
				{1, 0, 3},
			},
			expected: [][]int{
				{0, 0, 0},
			},
		},
		{
			name: "single column with zero",
			input: [][]int{
				{1},
				{0},
				{3},
			},
			expected: [][]int{
				{0},
				{0},
				{0},
			},
		},
		{
			name: "single element zero",
			input: [][]int{
				{0},
			},
			expected: [][]int{
				{0},
			},
		},
		{
			name: "single element non-zero",
			input: [][]int{
				{5},
			},
			expected: [][]int{
				{5},
			},
		},
		{
			name: "2x2 with one zero",
			input: [][]int{
				{1, 0},
				{3, 4},
			},
			expected: [][]int{
				{0, 0},
				{3, 0},
			},
		},
		{
			name: "entire row is zero",
			input: [][]int{
				{1, 2, 3},
				{0, 0, 0},
				{4, 5, 6},
			},
			expected: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "entire column is zero",
			input: [][]int{
				{1, 0, 3},
				{4, 0, 6},
				{7, 0, 9},
			},
			expected: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "rectangular matrix 2x4",
			input: [][]int{
				{1, 2, 3, 4},
				{5, 0, 7, 8},
			},
			expected: [][]int{
				{1, 0, 3, 4},
				{0, 0, 0, 0},
			},
		},
		{
			name: "rectangular matrix 4x2",
			input: [][]int{
				{1, 2},
				{3, 0},
				{5, 6},
				{7, 8},
			},
			expected: [][]int{
				{1, 0},
				{0, 0},
				{5, 0},
				{7, 0},
			},
		},
		{
			name: "diagonal zeros",
			input: [][]int{
				{0, 1, 1},
				{1, 0, 1},
				{1, 1, 0},
			},
			expected: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dsa_brushup.Matrix_set_zero(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Matrix_set_zero() failed\nInput:    %v\nGot:      %v\nExpected: %v", tt.input, result, tt.expected)
			}
		})
	}
}
