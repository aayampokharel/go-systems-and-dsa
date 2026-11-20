package main

import "testing"

type TestCase struct {
	name          string
	InputIntSlice []int
	OuptutInt     int
}

func TestLongestConsecutiveSequences(t *testing.T) {
	testCase := []TestCase{
		{
			"single digit",
			[]int{1},
			1,
		},
		{
			"randome 2 digits",
			[]int{1, 100},
			1,
		}, {
			"serial 2 numbers",
			[]int{1, 2},
			2,
		},

		{
			"duplicates",
			[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 12},
			1,
		},

		{
			"duplicates with sequence",
			[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 12, 1, 2, 3, 4, 5},
			5,
		},

		{
			"2 consecutive numbers",
			[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 12, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			2,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			got1 := longestConsecutive(tc.InputIntSlice)
			if got1 != tc.OuptutInt {
				t.Errorf("longestConsecutive(%v) = %v, want %v", tc.InputIntSlice, got1, tc.OuptutInt)

			}
			got2 := BestSolution(tc.InputIntSlice)
			if got2 != tc.OuptutInt {
				t.Errorf("BestSolution(%v) = %v, want %v", tc.InputIntSlice, got2, tc.OuptutInt)

			}
		})
	}
}
