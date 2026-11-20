package main

import (
	"fmt"
	"testing"
)

func generateBinaryString(i int) string {
	str := ""
	for range i {
		str += "a"
	}
	return str
}
func BenchmarkFindMaxConsecutiveOnes(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000, 100000}
	for _, val := range sizes {
		sizedSlice := generateBinaryString(val)
		b.Run(fmt.Sprintf("findMaxConsecutiveOnes() for size of::%v \n", val), func(b *testing.B) {
			for b.Loop() {
				_ = isPalindrome(sizedSlice)
			}
		})
	}
}

func BenchmarkBetterSolution(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000, 100000}
	for _, val := range sizes {
		sizedSlice := generateBinaryString(val)
		b.Run(fmt.Sprintf("BestSolution() for size: %v", val), func(b *testing.B) {
			for b.Loop() {
				_ = BestSolution(sizedSlice)
			}
		})
	}
}
