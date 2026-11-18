package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func generateEncodeStrings(count, length int) []string {
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

func BenchmarkEncode(b *testing.B) {
	s := &Solution{}
	sizes := []int{10, 100, 1000, 10000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size:%v", size), func(b *testing.B) {
			strs := generateEncodeStrings(size, 20)
			for b.Loop() {
				_ = s.Encode(strs)
			}
		})
	}
}

func BenchmarkDecode(b *testing.B) {
	s := &Solution{}
	sizes := []int{10, 100, 1000, 10000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size:%v", size), func(b *testing.B) {
			encoded := s.Encode(generateEncodeStrings(size, 20))
			for b.Loop() {
				_ = s.Decode(encoded)
			}
		})
	}
}
