package main

import (
	"testing"
)

type TestCase struct {
	name     string
	input    []int
	expected int
}

func TestConsecutiveOnes(t *testing.T) {
	test := []TestCase{
		{
			"single 0",
			[]int{0},
			0,
		},
		{
			"single 1",
			[]int{1},
			1,
		},
		{
			"all 0",
			[]int{0, 0, 0},
			0,
		},
		{
			"all 1",
			[]int{1, 1, 1},
			3,
		},
		{
			"mixed 0 and 1 with 0 at ends",
			[]int{0, 1, 1, 1, 1, 1, 1, 0},
			6,
		},
		{
			"0 at left",
			[]int{0, 1, 1, 1, 1, 1, 1},
			6,
		},
		{
			"0 at right",
			[]int{1, 1, 1, 1, 1, 1, 0},
			6,
		},
		{
			"0 in middle",
			[]int{1, 1, 1, 0, 1, 1},
			3,
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			got1 := findMaxConsecutiveOnes(tc.input)
			got2 := BestSolution(tc.input)
			if got2 != tc.expected {
				t.Errorf("BestSolution(%v) = %v, want %v", tc.input, got2, tc.expected)
			}
			if got1 != tc.expected {
				t.Errorf("findMaxConsecutiveOnes(%v) = %v, want %v", tc.input, got1, tc.expected)
			}
		})
	}
}
