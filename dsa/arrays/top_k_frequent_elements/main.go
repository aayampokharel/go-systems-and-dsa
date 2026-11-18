package main

func topKFrequent(nums []int, k int) []int {
	var returnInt []int
	m := make(map[int]int, len(nums))
	for _, val := range nums {
		m[val]++
	}
	var count int
	var selectedKey int

	if k > len(m) {
		k = len(m)
	}
	for range k {
		selectedKey = 0
		count = 0
		for key, val := range m {
			if val > count {
				selectedKey = key
				count = val
			}
		}
		returnInt = append(returnInt, selectedKey)
		delete(m, selectedKey)
	}

	return returnInt
}

func BestSolution(nums []int, k int) []int {
	var returningSlice []int
	frequencyBasedSlice := make([][]int, len(nums)+1)
	numsMap := make(map[int]int)
	for _, val := range nums {
		numsMap[val]++
	}

	for key, freq := range numsMap {
		frequencyBasedSlice[freq] = append(frequencyBasedSlice[freq], key)
	}

	for i := len(frequencyBasedSlice) - 1; i >= 0; i-- {
		if frequencyBasedSlice[i] != nil {
			returningSlice = append(returningSlice, frequencyBasedSlice[i]...)
			if len(returningSlice) >= k {
				return returningSlice[:k]
			}
		}
	}
	return returningSlice
}
