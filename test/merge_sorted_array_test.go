package dsa_brushup_test

import (
	"dsa-brushup"
	"reflect"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	tests := []struct {
		name     string
		arr1     []int
		arr2     []int
		expected []int
	}{
		{
			name:     "both arrays same size",
			arr1:     []int{1, 3, 5},
			arr2:     []int{2, 4, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "first array larger",
			arr1:     []int{1, 3, 5, 7, 9},
			arr2:     []int{2, 4},
			expected: []int{1, 2, 3, 4, 5, 7, 9},
		},
		{
			name:     "second array larger",
			arr1:     []int{1, 3},
			arr2:     []int{2, 4, 6, 8, 10},
			expected: []int{1, 2, 3, 4, 6, 8, 10},
		},
		{
			name:     "no overlap",
			arr1:     []int{1, 2, 3},
			arr2:     []int{4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "first array all smaller",
			arr1:     []int{1, 2, 3},
			arr2:     []int{7, 8, 9},
			expected: []int{1, 2, 3, 7, 8, 9},
		},
		{
			name:     "second array all smaller",
			arr1:     []int{7, 8, 9},
			arr2:     []int{1, 2, 3},
			expected: []int{1, 2, 3, 7, 8, 9},
		},
		{
			name:     "first array empty",
			arr1:     []int{},
			arr2:     []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "second array empty",
			arr1:     []int{1, 2, 3},
			arr2:     []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "both arrays empty",
			arr1:     []int{},
			arr2:     []int{},
			expected: []int{},
		},
		{
			name:     "single element each",
			arr1:     []int{1},
			arr2:     []int{2},
			expected: []int{1, 2},
		},
		{
			name:     "duplicate elements",
			arr1:     []int{1, 3, 5},
			arr2:     []int{1, 3, 5},
			expected: []int{1, 1, 3, 3, 5, 5},
		},
		{
			name:     "negative numbers",
			arr1:     []int{-5, -3, -1},
			arr2:     []int{-4, -2, 0},
			expected: []int{-5, -4, -3, -2, -1, 0},
		},
		{
			name:     "mixed positive and negative",
			arr1:     []int{-3, -1, 2, 4},
			arr2:     []int{-2, 0, 3, 5},
			expected: []int{-3, -2, -1, 0, 2, 3, 4, 5},
		},
		{
			name:     "with zeros",
			arr1:     []int{-2, 0, 2},
			arr2:     []int{-1, 0, 1},
			expected: []int{-2, -1, 0, 0, 1, 2},
		},
		{
			name:     "large arrays",
			arr1:     []int{1, 3, 5, 7, 9, 11, 13, 15},
			arr2:     []int{2, 4, 6, 8, 10, 12, 14, 16},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			name:     "interleaved elements",
			arr1:     []int{1, 4, 7, 10},
			arr2:     []int{2, 5, 8, 11},
			expected: []int{1, 2, 4, 5, 7, 8, 10, 11},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dsa_brushup.Merge_sorted_array(tt.arr1, tt.arr2)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Merge_sorted_array(%v, %v) = %v; want %v", tt.arr1, tt.arr2, result, tt.expected)
			}
		})
	}
}
