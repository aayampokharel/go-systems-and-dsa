package main

import (
	"fmt"
	"testing"
)

func generateEnglishLowerCasedString(i int) string {
	var letter string
	for key := range i {
		letter += string(byte('a' + (key % 26)))
	}
	return letter
}

func BenchmarkIsAnagram(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	for _, val := range sizes {
		b.Run(fmt.Sprintf("size of:%v", val), func(b *testing.B) {
			str := generateEnglishLowerCasedString(val)
			for b.Loop() {
				_ = isAnagram(str, str)
			}
		})
	}
}
func BenchmarkBestSolution(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	for _, val := range sizes {
		b.Run(fmt.Sprintf("size of for bestSolution :::%v", val), func(b *testing.B) {
			str := generateEnglishLowerCasedString(val)
			for b.Loop() {
				_ = BestSolution(str, str)
			}
		})
	}
}
