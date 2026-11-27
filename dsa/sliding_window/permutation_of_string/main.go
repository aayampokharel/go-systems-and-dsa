package main

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	lenOfS1 := len(s1)
	charMap := make(map[int32]int, len(s1))
	isSubstring := true

	for key := 0; key < len(s2); key++ {
		clear(charMap)
		for _, val := range s1 {
			charMap[val]++
		}
		isLoopExecuted := false
		isSubstring = true
		if _, ok1 := charMap[int32(s2[key])]; ok1 {
			rightPtr := key + lenOfS1 - 1
			for ; rightPtr < len(s2) && rightPtr >= key; rightPtr-- {
				isLoopExecuted = true
				if val2, ok2 := charMap[int32(s2[rightPtr])]; ok2 {
					charMap[int32(s2[rightPtr])]--
					if val2 <= 0 {
						isSubstring = false
						key = rightPtr
						break
					}
				} else {
					isSubstring = false
					key = rightPtr
					break
				}
			}
			if isSubstring && isLoopExecuted {
				return isSubstring
			}
			isLoopExecuted, isSubstring = false, true
		}

	}
	return false
}
