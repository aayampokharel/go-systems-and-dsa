package main

import "testing"

type TestCase struct {
	name          string
	inputIntSlice []int
	outputInt     int
}

func TestMaxProfit(t *testing.T) {
	testCase := []TestCase{
		{
			"empty slice",
			[]int{},
			0,
		},

		{
			"single element",
			[]int{1},
			0,
		},
		{
			"double element",
			[]int{1, 2},
			1,
		},
		{
			"double element(decreasing)",
			[]int{2, 1},
			0,
		},
		{
			"same element",
			[]int{1, 1},
			0,
		},
		{
			"multiple elements",
			[]int{2, 1, 5, 8, 9, 1, 10},
			9,
		},
		{
			"multiple elements",
			[]int{1, 8, 5, 8, 9, 9, 10},
			9,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			got := maxProfit(tc.inputIntSlice)
			if got != tc.outputInt {
				t.Errorf("maxProfit(): expected=%v,got=%v", tc.outputInt, got)
			}
		})
	}
}
