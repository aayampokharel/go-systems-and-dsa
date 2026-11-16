package main

import "fmt"

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}

func betterSolution(nums []int) []int { //betterSolution according to neet150
	n := len(nums)
	ans := make([]int, 2*n)
	for i := range n {
		ans[i] = nums[i]
		ans[i+n] = nums[i]
	}
	return ans
}

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 3}))
	fmt.Println(betterSolution([]int{1, 2, 3}))
}
