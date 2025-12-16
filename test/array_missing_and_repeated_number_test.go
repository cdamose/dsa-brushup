package dsa_brushup_test

import (
	"dsa-brushup"
	"testing"
)

func TestArrayMissingAndRepeatedNumber(t *testing.T) {
	tests := []struct {
		name             string
		input            []int
		expectedMissing  int
		expectedRepeated int
	}{
		{
			name:             "simple case",
			input:            []int{1, 2, 2, 4},
			expectedMissing:  3,
			expectedRepeated: 2,
		},
		{
			name:             "repeated at beginning",
			input:            []int{1, 1, 3, 4},
			expectedMissing:  2,
			expectedRepeated: 1,
		},
		{
			name:             "repeated at end",
			input:            []int{1, 2, 3, 3},
			expectedMissing:  4,
			expectedRepeated: 3,
		},
		{
			name:             "missing at beginning",
			input:            []int{2, 2, 3, 4},
			expectedMissing:  1,
			expectedRepeated: 2,
		},
		{
			name:             "missing at end",
			input:            []int{1, 2, 3, 3},
			expectedMissing:  4,
			expectedRepeated: 3,
		},
		{
			name:             "two elements",
			input:            []int{1, 1},
			expectedMissing:  2,
			expectedRepeated: 1,
		},
		{
			name:             "three elements",
			input:            []int{1, 2, 2},
			expectedMissing:  3,
			expectedRepeated: 2,
		},
		{
			name:             "larger array",
			input:            []int{1, 2, 3, 4, 5, 5, 7},
			expectedMissing:  6,
			expectedRepeated: 5,
		},
		{
			name:             "unsorted array",
			input:            []int{3, 1, 2, 5, 3},
			expectedMissing:  4,
			expectedRepeated: 3,
		},
		{
			name:             "repeated in middle",
			input:            []int{1, 2, 3, 3, 5},
			expectedMissing:  4,
			expectedRepeated: 3,
		},
		{
			name:             "missing in middle",
			input:            []int{1, 2, 4, 4, 5},
			expectedMissing:  3,
			expectedRepeated: 4,
		},
		{
			name:             "array 1 to 5",
			input:            []int{1, 2, 3, 4, 4},
			expectedMissing:  5,
			expectedRepeated: 4,
		},
		{
			name:             "array with 1 repeated",
			input:            []int{1, 1, 3, 4, 5},
			expectedMissing:  2,
			expectedRepeated: 1,
		},
		{
			name:             "array with last repeated",
			input:            []int{1, 2, 3, 4, 5, 6, 7, 7, 9, 10},
			expectedMissing:  8,
			expectedRepeated: 7,
		},
		{
			name:             "consecutive missing and repeated",
			input:            []int{1, 2, 3, 3, 5},
			expectedMissing:  4,
			expectedRepeated: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			missing, repeated := dsa_brushup.Array_missing_and_repeated_number(tt.input)
			
			if missing != tt.expectedMissing {
				t.Errorf("Array_missing_and_repeated_number(%v) missing = %d; want %d", 
					tt.input, missing, tt.expectedMissing)
			}
			
			if repeated != tt.expectedRepeated {
				t.Errorf("Array_missing_and_repeated_number(%v) repeated = %d; want %d", 
					tt.input, repeated, tt.expectedRepeated)
			}
		})
	}
}
