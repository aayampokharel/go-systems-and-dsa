package main

func lengthOfLongestSubstring(s string) int { //abcd(4)..............d(4)c(1)b(2)a(3)
	runeMap := make(map[int32]int, len(s))
	returnIntLength := 0
	runeSubstring := make([]int32, 0, len(s))
	for _, val := range s {
		if _, ok := runeMap[val]; !ok {
			runeMap[val]++

		} else {
			for key, val1 := range runeSubstring {
				if val1 == val {
					runeSubstring = runeSubstring[key+1:]
					break
				}

			}
		}
		runeSubstring = append(runeSubstring, val)

		if len(runeSubstring) > returnIntLength {
			returnIntLength = len(runeSubstring)
		}

	}
	return returnIntLength
}
