package main

import (
	"fmt"
	"testing"
)

func generateBinarySlice(i int) []int {
	var slice []int
	for key := range i {
		if key%2 == 0 {
			slice = append(slice, 0)
		} else {
			slice = append(slice, 1)
		}
	}
	return slice
}
func BenchmarkFindMaxConsecutiveOnes(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000, 100000}
	for _, val := range sizes {
		sizedSlice := generateBinarySlice(val)
		b.Run(fmt.Sprintf("longestconsecutiveSequence() for size of::%v \n", val), func(b *testing.B) {
			for b.Loop() {
				_ = longestConsecutive(sizedSlice)
			}
		})
	}
}

func BenchmarkBetterSolution(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000, 100000}
	for _, val := range sizes {
		sizedSlice := generateBinarySlice(val)
		b.Run(fmt.Sprintf("bestSolution() for size: %v", val), func(b *testing.B) {
			for b.Loop() {
				_ = BestSolution(sizedSlice)
			}
		})
	}
}
