package main

import (
	"fmt"
	"testing"
)

func generateSlice(i int) []int {
	var slice []int
	for range i {
		slice = append(slice, i)
	}
	return slice
}
func BenchmarkGetConcatenation(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	for _, val := range sizes {

		b.Run(fmt.Sprintf("size of:%v", val), func(b *testing.B) {
			nums := generateSlice(val)
			for b.Loop() {
				_ = getConcatenation(nums)
			}
		})

	}
}

func BenchmarkBetterSolution(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	for _, val := range sizes {

		b.Run(fmt.Sprintf("size of:%v", val), func(b *testing.B) {
			nums := generateSlice(val)
			for b.Loop() {
				_ = betterSolution(nums)
			}
		})

	}
}
