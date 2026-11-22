package main

import "math"

func maxArea(heights []int) int {

	highestArea := 0
	leftPtr := 0
	rightPtr := len(heights) - 1

	for leftPtr < rightPtr {
		minHeight := min(heights[leftPtr], heights[rightPtr])
		width := math.Abs(float64(rightPtr) - float64(leftPtr))
		area := minHeight * int(width)
		if area > highestArea {
			highestArea = area
		}
		if heights[leftPtr] == minHeight {
			leftPtr++
		} else {
			rightPtr--
		}
	}
	return highestArea
}
