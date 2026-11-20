package main

func longestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	mNums := make(map[int]int, len(nums))
	var count int = 1
	var tempCount int = 1
	for _, val := range nums {
		mNums[val]++
	}

	for _, val := range nums { //203
		if _, ok := mNums[val-1]; !ok { //4(1,300,400,500)

			for i := 0; i < len(nums); i++ { //200
				if _, ok2 := mNums[val+i+1]; ok2 {
					tempCount++
				} else {
					if tempCount > count {
						count = tempCount
					}
					tempCount = 1
					break
				}

			}
		}
	}
	return count

}

func BestSolution(nums []int) int {
	numSet := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	longest := 0

	for num := range numSet {
		// Only start counting from the beginning of a sequence.
		if _, exists := numSet[num-1]; !exists {
			current := num
			count := 1

			for {
				if _, exists := numSet[current+1]; !exists {
					break
				}

				current++
				count++
			}

			if count > longest {
				longest = count
			}
		}
	}

	return longest
}
