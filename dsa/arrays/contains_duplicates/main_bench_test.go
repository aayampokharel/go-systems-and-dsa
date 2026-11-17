package main

import (
	"fmt"
	"testing"
)

func generateSlice(i int) []int {
	slice := make([]int, 0, i)
	for key := range i {
		slice = append(slice, key)
	}
	return slice
}
func BenchmarkContainsDuplicate(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	for _, val := range sizes {

		b.Run(fmt.Sprintf("size of for generateSlice:%v", val), func(b *testing.B) {
			nums := generateSlice(val)
			for b.Loop() {
				_ = hasDuplicate(nums)
			}
		})

	}
}

func BenchmarkBestSolution(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	for _, val := range sizes {

		b.Run(fmt.Sprintf("size of for BestSolution:%v", val), func(b *testing.B) {
			nums := generateSlice(val)
			for b.Loop() {
				_ = BestSolution(nums)
			}
		})

	}
}
