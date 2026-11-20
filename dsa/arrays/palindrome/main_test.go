package main

import "testing"

type TestCase struct {
	name        string
	inputString string
	output      bool
}

func TestIsPalindrome(t *testing.T) {
	testCase := []TestCase{
		{
			"single character",
			"1",
			true,
		},
		{
			"double character",
			"12",
			false,
		},
		{
			"double character",
			"11",
			true,
		},
		{
			"empty string",
			"",
			true,
		},
		{
			"valid palindrome",
			"racecar",
			true,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			got1 := isPalindrome(tc.inputString)
			if got1 != tc.output {
				t.Errorf("findMaxConsecutiveOnes(%v) = %v, want %v", tc.inputString, got1, tc.output)
			}
			got2 := BestSolution(tc.inputString)
			if got2 != tc.output {
				t.Errorf("BetterSolution(%v) = %v, want %v", tc.inputString, got2, tc.output)
			}
		})
	}

}
