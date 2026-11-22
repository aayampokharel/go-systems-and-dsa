package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	name          string
	inputSlice    []int
	output2DSlice [][]int
}

func TestThreeSum(t *testing.T) {
	testCases := []TestCase{
		{"3 elements",
			[]int{0, 0, 0},
			[][]int{{0, 0, 0}},
		},
		{
			"4 elements(should avoid duplicates)",
			[]int{0, 0, 0, 0},
			[][]int{{0, 0, 0}},
		},
		{
			"multiple combinations",
			[]int{-3, -2, -1, 0, 1, 2, 3},
			[][]int{{-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}},
		},
		{
			"no combinations",
			[]int{-3, -2, 0},
			[][]int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := append([]int(nil), tc.inputSlice...)

			got := threeSum(input)

			if !reflect.DeepEqual(got, tc.output2DSlice) {
				t.Errorf(
					"threeSum() = %v, want %v",
					got,
					tc.output2DSlice,
				)
			}
			got = BetterSolution(input)

			if !reflect.DeepEqual(got, tc.output2DSlice) {
				t.Errorf(
					"BetterSolution() = %v, want %v",
					got,
					tc.output2DSlice,
				)
			}
		})
	}
}
