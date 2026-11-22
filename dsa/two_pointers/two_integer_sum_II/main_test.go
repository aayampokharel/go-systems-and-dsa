package main

import (
	"slices"
	"testing"
)

type TestCase struct {
	name                string
	inputSortedSliceInt []int
	inputTarget         int
	Indexoutcome        []int
}

func TestTwoSum(t *testing.T) {
	tc := []TestCase{
		{
			"single element(doesnot satisfy sum since only 1 element)",
			[]int{1},
			1,
			[]int{},
		},
		{
			"two elements",
			[]int{1, 2},
			3,
			[]int{1, 2},
		},
		{
			"two elements/wrong target",
			[]int{1, 2},
			5,
			[]int{},
		},
		{
			"odd number ",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			11,
			[]int{1, 10},
		},
		{
			"even number ",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			11,
			[]int{1, 10},
		},
	}

	for _, val := range tc {
		t.Run(val.name, func(t *testing.T) {
			got := twoSum(val.inputSortedSliceInt, val.inputTarget)
			if !slices.Equal(got, val.Indexoutcome) {
				t.Errorf("twoSum(%v,%v)=> expected:%v::::got:%v", val.inputSortedSliceInt, val.inputTarget, val.Indexoutcome, got)
			}
		})
	}
}
