package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func generateNums(size int) []int {
	r := rand.New(rand.NewSource(42))
	nums := make([]int, size)
	for i := range nums {
		nums[i] = r.Intn(size/10+1) - 5
	}
	return nums
}

func BenchmarkTopKFrequent(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size:%v", size), func(b *testing.B) {
			nums := generateNums(size)
			k := size / 10
			if k < 1 {
				k = 1
			}
			for b.Loop() {
				_ = topKFrequent(nums, k)
			}
		})
	}
}

func BenchmarkBestSolution(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size:%v", size), func(b *testing.B) {
			nums := generateNums(size)
			k := size / 10
			if k < 1 {
				k = 1
			}
			for b.Loop() {
				_ = BestSolution(nums, k)
			}
		})
	}
}
