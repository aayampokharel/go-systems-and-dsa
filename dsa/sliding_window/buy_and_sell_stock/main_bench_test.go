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
func BenchmarkMaxProfit(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	for _, val := range sizes {

		b.Run(fmt.Sprintf("size of for generateSlice:%v", val), func(b *testing.B) {
			nums := generateSlice(val)
			for b.Loop() {
				_ = maxProfit(nums)
			}
		})

	}
}

func BenchmarkBetterSolution(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	for _, val := range sizes {

		b.Run(fmt.Sprintf("size of for BetterSolution:%v", val), func(b *testing.B) {
			nums := generateSlice(val)
			for b.Loop() {
				_ = BetterSolution(nums)
			}
		})

	}
}
