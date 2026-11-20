package main

import "strings"

func keepAlphanumeric(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	for _, r := range s {
		if (r >= '0' && r <= '9') ||
			(r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func isPalindrome(s string) bool { //raccar, 1rac + car1
	s = strings.ToLower(keepAlphanumeric(s))
	strByte := []byte(s)
	mid := len(s) / 2
	if mid%2 != 0 && mid > 1 {
		s = string(append(strByte[:mid], strByte[mid+1:]...))
	}
	for i := 0; i < mid; i++ {
		if string(s[i]) != string(s[len(s)-1-i]) {
			return false
		}

	}

	return true
}

func BestSolution(s string) bool {
	s = strings.ToLower(keepAlphanumeric(s))

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
