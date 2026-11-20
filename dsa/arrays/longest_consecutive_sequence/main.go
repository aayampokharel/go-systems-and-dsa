package main

func longestConsecutive(nums []int) int {
	// nums:=slices.sort(nums)
	// var tempSlice []int
	// var returnSlice []int
	//
	//	for i:=0;i<len(nums)-1;i++{
	//		if nums[i]+1==nums[i+1]{
	//	       tempSlice=append(tempSlice,nums[i])
	//		}
	//	}
	if len(nums) == 0 {
		return 0
	}
	mNums := make(map[int]int, len(nums))
	var count int = 1
	var tempCount int = 1
	for _, val := range nums {
		mNums[val]++
	}

	for _, val := range nums { //n
		if _, ok := mNums[val-1]; !ok {  //n

			for i := 0; i < len(nums); i++ {
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
