/*
Given a sorted array of distinct integers and a target value,
return the index if the target is found. If not, return the index where
it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.


Example 1:
Input: nums = [1,3,5,6], target = 5
Output: 2

Example 2:
Input: nums = [1,3,5,6], target = 2
Output: 1

Example 3:
Input: nums = [1,3,5,6], target = 7
Output: 4
*/

package Alghorithm

// FindInsertPosition
// Implements binary search algorithm in a recursive way to find an index of value in sorted array
// or to find an index in array to insert a value
func FindInsertPosition(nums []int, target int) int {
	var binarySearchRecursive func(arr []int, valueToFind, startIdx int) int

	binarySearchRecursive = func(arr []int, valueToFind, startIdx int) int {
		medianIdx := len(arr) / 2
		medianValue := arr[medianIdx]
		if len(arr) == 1 {
			if medianValue == valueToFind || valueToFind < medianValue {
				return startIdx
			}
			return startIdx + 1
		}
		if valueToFind > medianValue {
			return binarySearchRecursive(arr[medianIdx:], valueToFind, medianIdx+startIdx)
		}
		if valueToFind < medianValue {
			return binarySearchRecursive(arr[:medianIdx], valueToFind, startIdx)
		}
		return medianIdx + startIdx
	}

	return binarySearchRecursive(nums, target, 0)
}
