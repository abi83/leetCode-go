/*
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
*/
package Alghorithm

// RotateArrayMemoryEffective
// implementation "on place" without using additional memory
func RotateArrayMemoryEffective(nums []int, k int) []int {
	for rotationNum := 0; rotationNum < k; rotationNum++ {
		for i := 0; i < len(nums)-1; i++ {
			nums[len(nums)-1], nums[i] = nums[i], nums[len(nums)-1]
		}
	}
	return nums
}

// RotateArrayCpuEffective
// implementation without looping k times
func RotateArrayCpuEffective(nums []int, k int) {
	slide := (20*len(nums) + len(nums) - k) % len(nums)
	copy(nums, append(nums[slide:], nums[:slide]...))
}
