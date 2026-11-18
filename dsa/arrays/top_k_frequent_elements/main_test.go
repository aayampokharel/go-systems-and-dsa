package main

import (
	"reflect"
	"sort"
	"testing"
)

func equalIntSets(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	ac := append([]int(nil), a...)
	bc := append([]int(nil), b...)
	sort.Ints(ac)
	sort.Ints(bc)
	return reflect.DeepEqual(ac, bc)
}

type topKTestCase struct {
	name string
	nums []int
	k    int
	want []int
}

func TestTopKFrequent(t *testing.T) {
	tests := []topKTestCase{
		{name: "example", nums: []int{1, 2, 2, 3, 3, 3}, k: 2, want: []int{2, 3}},
		{name: "single value", nums: []int{7, 7}, k: 1, want: []int{7}},
		{name: "one element", nums: []int{1}, k: 1, want: []int{1}},
		{name: "all unique", nums: []int{1, 2, 3}, k: 3, want: []int{1, 2, 3}},
		{name: "negative values", nums: []int{-1, -1, -2, -2, -2}, k: 1, want: []int{-2}},
		{name: "including zero", nums: []int{0, 0, 0, 1, 1, 2}, k: 2, want: []int{0, 1}},
		{name: "k equals one", nums: []int{4, 4, 4, 5, 5, 6}, k: 1, want: []int{4}},
		{name: "mixed frequencies", nums: []int{9, 9, 9, 9, 1, 1, 1, 2, 2, 3}, k: 3, want: []int{9, 1, 2}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := topKFrequent(tc.nums, tc.k)
			if !equalIntSets(got, tc.want) {
				t.Errorf("topKFrequent(%v, %d) = %v, want %v", tc.nums, tc.k, got, tc.want)
			}

			gotBest := BestSolution(tc.nums, tc.k)
			if !equalIntSets(gotBest, tc.want) {
				t.Errorf("BestSolution(%v, %d) = %v, want %v", tc.nums, tc.k, gotBest, tc.want)
			}
		})
	}
}
