package main

import "testing"

type TestCase struct {
	name     string
	input    []int
	expected bool
}

func TestContainsDuplicate(t *testing.T) {
	tests := []TestCase{
		{
			name:     "single 0",
			input:    []int{0},
			expected: false,
		},
		{
			name:     "single 1",
			input:    []int{1},
			expected: false,
		},
		{
			name:     "all 0",
			input:    []int{0, 0, 0},
			expected: true,
		},
		{
			name:     "all 1",
			input:    []int{1, 1, 1},
			expected: true,
		},
		{
			name:     "no duplicate",
			input:    []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "duplicate of length>1",
			input:    []int{1, 2, 3, 4, 5, 6, 1},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := hasDuplicate(tc.input)
			if got != tc.expected {
				t.Errorf("hasDuplicate(%v) = %v, want %v", tc.input, got, tc.expected)
			}
			bestGot := BestSolution(tc.input)
			if bestGot != tc.expected {
				t.Errorf("BestSolution(%v) = %v, want %v", tc.input, got, tc.expected)
			}

		})
	}
}
