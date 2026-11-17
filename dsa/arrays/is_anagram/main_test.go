package main

import "testing"

type TestCase struct {
	name     string
	s        string
	t        string
	expected bool
}

func TestIsAnagram(t *testing.T) {
	tc := []TestCase{
		{
			"empty",
			"",
			"",
			true,
		},
		{
			"single letter",
			"a",
			"a",
			true,
		},
		{
			"2 letters",
			"ab",
			"ba",
			true,
		},
		{
			"3 letters",
			"abc",
			"bca",
			true,
		},
		{
			"non anagram",
			"abc",
			"bcd",
			false,
		},
		{
			"non anagram",
			"aa",
			"bb",
			false,
		},
	}

	for _, val := range tc {
		t.Run(val.name, func(t *testing.T) {
			got := isAnagram(val.s, val.t)
			if got != val.expected {
				t.Errorf("isAnagram(%v,%v)=%v,want %v", val.s, val.t, got, val.expected)
			}
			got1 := BestSolution(val.s, val.t)
			if got1 != val.expected {
				t.Errorf("BestSolution(%v,%v)=%v,want %v", val.s, val.t, got1, val.expected)
			}
		})
	}
}
