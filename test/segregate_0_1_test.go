package dsa_brushup_test

import (
	"dsa-brushup"
	"testing"
)

func TestSegregate01(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "mixed 0s and 1s",
			input: []int{0, 1, 0, 1, 1, 0},
		},
		{
			name:  "all zeros",
			input: []int{0, 0, 0, 0},
		},
		{
			name:  "all ones",
			input: []int{1, 1, 1, 1},
		},
		{
			name:  "already segregated",
			input: []int{0, 0, 1, 1},
		},
		{
			name:  "reverse segregated",
			input: []int{1, 1, 0, 0},
		},
		{
			name:  "single zero",
			input: []int{0},
		},
		{
			name:  "single one",
			input: []int{1},
		},
		{
			name:  "two elements 0,1",
			input: []int{0, 1},
		},
		{
			name:  "two elements 1,0",
			input: []int{1, 0},
		},
		{
			name:  "alternating pattern",
			input: []int{0, 1, 0, 1, 0, 1},
		},
		{
			name:  "mostly zeros",
			input: []int{0, 0, 0, 1, 0, 0},
		},
		{
			name:  "mostly ones",
			input: []int{1, 1, 1, 0, 1, 1},
		},
		{
			name:  "zeros at beginning",
			input: []int{0, 0, 0, 1, 1, 1},
		},
		{
			name:  "ones at beginning",
			input: []int{1, 1, 1, 0, 0, 0},
		},
		{
			name:  "large array",
			input: []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Count zeros in input
			zeroCount := 0
			for _, val := range tt.input {
				if val == 0 {
					zeroCount++
				}
			}

			result := dsa_brushup.Segregate_0_1(tt.input)

			// Verify all zeros are at the beginning
			for i := 0; i < zeroCount; i++ {
				if result[i] != 0 {
					t.Errorf("Segregate_0_1() position %d should be 0, got %d", i, result[i])
				}
			}

			// Verify all ones are at the end
			for i := zeroCount; i < len(result); i++ {
				if result[i] != 1 {
					t.Errorf("Segregate_0_1() position %d should be 1, got %d", i, result[i])
				}
			}

			// Verify length is preserved
			if len(result) != len(tt.input) {
				t.Errorf("Segregate_0_1() length = %d; want %d", len(result), len(tt.input))
			}
		})
	}
}
