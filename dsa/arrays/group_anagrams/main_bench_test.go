package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func generateLowercaseStrings(count, length int) []string {
	r := rand.New(rand.NewSource(42))
	out := make([]string, count)
	buf := make([]byte, length)
	for i := range out {
		for j := range buf {
			buf[j] = byte('a' + r.Intn(26))
		}
		out[i] = string(buf)
	}
	return out
}

func BenchmarkGroupAnagrams(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size:%v", size), func(b *testing.B) {
			strs := generateLowercaseStrings(size, 10)
			for b.Loop() {
				_ = groupAnagrams(strs)
			}
		})
	}
}

func BenchmarkBetterApproach(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size:%v", size), func(b *testing.B) {
			strs := generateLowercaseStrings(size, 10)
			for b.Loop() {
				_ = BetterApproach(strs)
			}
		})
	}
}
