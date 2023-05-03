package Alghorithm

import (
	"math"
)

func SortedSquares(nums []int) []int {
	var leftIndex, rightIndex = 0, len(nums) - 1
	var i, sqr = rightIndex, 0
	res := make([]int, len(nums))

	for leftIndex <= rightIndex {
		leftAbs, rightAbs := int(math.Abs(float64(nums[leftIndex]))), int(math.Abs(float64(nums[rightIndex])))

		if leftAbs >= rightAbs {
			sqr = nums[leftIndex] * nums[leftIndex]
			leftIndex++
		}
		if rightAbs >= leftAbs {
			sqr = nums[rightIndex] * nums[rightIndex]
			rightIndex--
		}
		res[i] = sqr
		i--
	}
	return res
}
