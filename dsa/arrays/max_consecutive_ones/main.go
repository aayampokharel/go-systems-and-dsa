package main

func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	for key := 0; key < len(nums); key++ {
		if nums[key] == 1 {
			currentCount := 1
			for i := key + 1; i < len(nums); i++ {
				if nums[i] == 1 {
					currentCount++
				} else {
					key = i
					break
				}
			}
			if currentCount > maxCount {
				maxCount = currentCount
			}
		}
	}
	return maxCount
}

func BestSolution(nums []int) int {
	maxCount := 0
	currentCount := 0
	for _, val := range nums {
		if val == 1 {
			currentCount++
			if maxCount < currentCount {
				maxCount = currentCount
			}
		} else {
			currentCount = 0
		}
	}
	return maxCount
}
