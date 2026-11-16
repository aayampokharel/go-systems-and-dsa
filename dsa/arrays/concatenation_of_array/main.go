package main

import "fmt"

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}

func betterSolution(nums []int) []int {
	l := len(nums)
	res := make([]int, l*2)
	for i := range l {
		res[i] = nums[i]
		res[l+i] = nums[i]
	}
	return res
}

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 3}))
	fmt.Println(betterSolution([]int{1, 2, 3}))
}
