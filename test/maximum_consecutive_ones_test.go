package dsa_brushup_test

import (
	"dsa-brushup"
	"testing"
)

func TestMaximumConsecutiveOneInArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "consecutive ones at beginning",
			input:    []int{1, 1, 1, 0, 1, 1},
			expected: 3,
		},
		{
			name:     "consecutive ones at end",
			input:    []int{1, 0, 1, 1, 1, 1},
			expected: 4,
		},
		{
			name:     "consecutive ones in middle",
			input:    []int{0, 1, 1, 1, 1, 0},
			expected: 4,
		},
		{
			name:     "all ones",
			input:    []int{1, 1, 1, 1, 1},
			expected: 5,
		},
		{
			name:     "no ones",
			input:    []int{0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "single one",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "single zero",
			input:    []int{0},
			expected: 0,
		},
		{
			name:     "alternating ones and zeros",
			input:    []int{1, 0, 1, 0, 1, 0},
			expected: 1,
		},
		{
			name:     "multiple sequences same length",
			input:    []int{1, 1, 0, 1, 1, 0, 1, 1},
			expected: 2,
		},
		{
			name:     "two ones separated by zero",
			input:    []int{1, 1, 0, 1, 1, 1},
			expected: 3,
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "long sequence at start",
			input:    []int{1, 1, 1, 1, 1, 1, 0, 1, 1},
			expected: 6,
		},
		{
			name:     "long sequence at end",
			input:    []int{1, 1, 0, 1, 1, 1, 1, 1, 1},
			expected: 6,
		},
		{
			name:     "multiple zeros between ones",
			input:    []int{1, 1, 0, 0, 0, 1, 1, 1},
			expected: 3,
		},
		{
			name:     "two elements both ones",
			input:    []int{1, 1},
			expected: 2,
		},
		{
			name:     "two elements both zeros",
			input:    []int{0, 0},
			expected: 0,
		},
		{
			name:     "two elements one and zero",
			input:    []int{1, 0},
			expected: 1,
		},
		{
			name:     "large array with max in middle",
			input:    []int{1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1},
			expected: 7,
		},
		{
			name:     "only one occurrence of one",
			input:    []int{0, 0, 1, 0, 0},
			expected: 1,
		},
		{
			name:     "multiple single ones",
			input:    []int{1, 0, 1, 0, 1, 0, 1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dsa_brushup.Maximum_consecutive_one_in_array(tt.input)
			if result != tt.expected {
				t.Errorf("Maximum_consecutive_one_in_array(%v) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
