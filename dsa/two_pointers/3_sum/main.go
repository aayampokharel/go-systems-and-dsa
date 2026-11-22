package main

import "slices"

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	var return2DSlice [][]int = [][]int{}
	var return2DSliceNonDuplicate [][]int = [][]int{}

	leftPtr := 0
	rightPtr := len(nums) - 1
	slices.Sort(nums)
	var sum int = 0

	for i := 0; i < len(nums)-1; i++ {
		leftPtr = i + 1
		rightPtr = len(nums) - 1
		for leftPtr < rightPtr {
			sum = nums[i] + nums[leftPtr] + nums[rightPtr]
			if sum == 0 {
				return2DSlice = append(return2DSlice, []int{nums[i], nums[leftPtr], nums[rightPtr]})
				leftPtr++
				rightPtr--
			} else if sum < 0 {
				leftPtr++
			} else {
				rightPtr--
			}
		}
	}

	var isDuplicate bool
	for _, val1 := range return2DSlice {
		if len(return2DSliceNonDuplicate) == 0 {
			return2DSliceNonDuplicate = append(return2DSliceNonDuplicate, val1)
		}
		for _, val2 := range return2DSliceNonDuplicate {
			if slices.Equal(val2, val1) {
				isDuplicate = true
			}
		}
		if !isDuplicate {
			return2DSliceNonDuplicate = append(return2DSliceNonDuplicate, val1)
		}
		isDuplicate = false
	}
	return return2DSliceNonDuplicate
}

func BetterSolution(nums []int) [][]int {
	result := [][]int{}

	if len(nums) < 3 {
		return result
	}

	slices.Sort(nums)

	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicate first elements.
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		leftPtr := i + 1
		rightPtr := len(nums) - 1

		for leftPtr < rightPtr {
			sum := nums[i] + nums[leftPtr] + nums[rightPtr]

			switch {
			case sum == 0:
				result = append(result, []int{
					nums[i],
					nums[leftPtr],
					nums[rightPtr],
				})

				leftPtr++
				rightPtr--

				// Skip duplicate left values.
				for leftPtr < rightPtr &&
					nums[leftPtr] == nums[leftPtr-1] {
					leftPtr++
				}

				// Skip duplicate right values.
				for leftPtr < rightPtr &&
					nums[rightPtr] == nums[rightPtr+1] {
					rightPtr--
				}

			case sum < 0:
				leftPtr++

			default:
				rightPtr--
			}
		}
	}
	return result
}
