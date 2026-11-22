package main

import "testing"

type TestCase struct {
	name       string
	inputSlice []int
	expected   int
}

func TestMaxArea(t *testing.T) {
	testCases := []TestCase{
		{
			name:       "basic example",
			inputSlice: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected:   49,
		},
		{
			name:       "two elements",
			inputSlice: []int{1, 1},
			expected:   1,
		},
		{
			name:       "increasing heights",
			inputSlice: []int{1, 2, 3, 4, 5},
			expected:   6,
		},
		{
			name:       "decreasing heights",
			inputSlice: []int{5, 4, 3, 2, 1},
			expected:   6,
		},
		{
			name:       "all equal heights",
			inputSlice: []int{5, 5, 5, 5},
			expected:   15,
		},
		{
			name:       "one height is zero",
			inputSlice: []int{0, 5, 0, 4},
			expected:   8,
		},
		{
			name:       "maximum area in middle",
			inputSlice: []int{2, 3, 10, 5, 7, 8, 9},
			expected:   36,
		},
		{
			name:       "empty input",
			inputSlice: []int{},
			expected:   0,
		},
		{
			name:       "single element",
			inputSlice: []int{10},
			expected:   0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := maxArea(tc.inputSlice)

			if got != tc.expected {
				t.Errorf(
					"maxArea(%v) = %d, want %d",
					tc.inputSlice,
					got,
					tc.expected,
				)
			}
		})
	}
}
