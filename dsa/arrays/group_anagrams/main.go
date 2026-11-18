package main

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	var returnStr [][]string
	m := make(map[string][]string, len(strs))

	for _, val := range strs {
		runes := []rune(val)
		slices.Sort(runes)
		m[string(runes)] = append(m[string(runes)], val)
	}

	for _, val := range m {
		returnStr = append(returnStr, val)
	}
	return returnStr
}

func BetterApproach(strs []string) [][]string {
	// Use a fixed-size array [26]byte as the map key because arrays are comparable in Go
	m := make(map[[26]byte][]string)

	for _, s := range strs {
		var count [26]byte
		for i := 0; i < len(s); i++ {
			count[s[i]-'a']++
		}
		m[count] = append(m[count], s)
	}

	// Preallocate the result slice for better performance
	result := make([][]string, 0, len(m))
	for _, group := range m {
		result = append(result, group)
	}

	return result
}
