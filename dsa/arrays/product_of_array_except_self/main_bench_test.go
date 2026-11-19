package main

import (
	"fmt"
	"testing"
)

func ConcatenateSameSlice(size int) []int {
	var duplicate []int = []int{1, 2, 3, 4}
	slice := make([]int, 0, size*len(duplicate))
	for range size {
		slice = append(slice, duplicate...)
	}
	return slice
}

func BenchmarkProductOfArrayExceptSelf(b *testing.B) {
	sizes := []int{10, 100, 1000, 5000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size of(for my soln):%v", size), func(b *testing.B) {
			for b.Loop() {
				_ = productExceptSelf(ConcatenateSameSlice(size))
			}
		})
	}

}
func BenchmarkBestSolution(b *testing.B) {
	sizes := []int{10, 100, 1000, 5000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size of(BEST SOLUTION):%v", size), func(b *testing.B) {
			for b.Loop() {
				_ = BestSolution(ConcatenateSameSlice(size))
			}
		})
	}

}
