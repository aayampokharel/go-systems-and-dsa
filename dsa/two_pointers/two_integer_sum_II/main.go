package main

func twoSum(numbers []int, target int) []int {

	leftPtr := 0
	rightPtr := len(numbers) - 1

	for leftPtr < rightPtr {
		sum := numbers[leftPtr] + numbers[rightPtr]
		if sum == target {
			return []int{leftPtr + 1, rightPtr + 1}
		} else if sum < target {
			leftPtr++
		} else {
			rightPtr--
		}

	}
	return []int{}
}
