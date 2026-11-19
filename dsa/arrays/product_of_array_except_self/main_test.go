package main

import (
	"slices"
	"testing"
)

type TestCase struct {
	name     string
	input    []int
	expected []int
}

func TestProductOfArrayExceptSelf(t *testing.T) {
	tc := []TestCase{ //input will have atleast 2 elements.
		{
			"least possible 2 elements",
			[]int{1, 2},
			[]int{2, 1},
		},
		{
			"least possible 2 same elements",
			[]int{1, 1},
			[]int{1, 1},
		},
		{
			"least possible 2 elements(both 0)",
			[]int{0, 0},
			[]int{0, 0},
		},
		{
			"length of 5(odd length)",
			[]int{1, 2, 3, 1, 1},
			[]int{6, 3, 2, 6, 6},
		},
		{
			"length of 5(even length)",
			[]int{1, 2, 3, 1, 1, 1},
			[]int{6, 3, 2, 6, 6, 6},
		},
		{
			"negative number ",
			[]int{1, 2, 3, 1, 1, -1},
			[]int{-6, -3, -2, -6, -6, 6},
		},
		{
			"two negative number ",
			[]int{-1, 2, 3, 1, 1, -1},
			[]int{-6, 3, 2, 6, 6, -6},
		},
	}

	for _, val := range tc {
		t.Run(val.name, func(t *testing.T) {
			got := productExceptSelf(val.input)
			if !slices.Equal(got, val.expected) {
				t.Errorf("productExceptSelf(%v) = %v, want %v", val.input, got, val.expected)
			}

			got2 := BestSolution(val.input)
			if !slices.Equal(got2, val.expected) {
				t.Errorf("BestSolution(%v) = %v, want %v", val.input, got2, val.expected)
			}
		})
	}
}
