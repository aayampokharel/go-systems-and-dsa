package main

import (
	"slices"
	"testing"
)

type TestCase struct {
	name     string
	nums     []int
	expected []int
}

func TestGetConcatenation(t *testing.T) {
	tests := []TestCase{
		{
			name:     "length of 1",
			nums:     []int{1},
			expected: []int{1, 1},
		},
		{
			name:     "length of 3(odd length)",
			nums:     []int{1, 2, 3},
			expected: []int{1, 2, 3, 1, 2, 3},
		},
		{
			name:     "empty slice",
			nums:     []int{},
			expected: []int{},
		},
		{
			name:     "length of 4 (even length)",
			nums:     []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4, 1, 2, 3, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := getConcatenation(tc.nums)

			if !slices.Equal(got, tc.expected) {
				t.Errorf("getConcatenation(%v) = %v, want %v", tc.nums, got, tc.expected)
			}
		})

	}
}
