package Alghorithm

func BinarySearch(nums []int, target int) int {
	var idx = len(nums) / 2
	var step = len(nums) / 2

	if nums[0] == target {
		return 0
	}
	if nums[len(nums)-1] == target {
		return len(nums) - 1
	}

	for nums[idx] != target &&
		idx >= 0 &&
		idx < len(nums)-1 {
		if step >= 2 {
			step = step / 2
		}
		if nums[idx] < target {
			idx = idx + step
		} else if nums[idx] > target {
			idx = idx - step
		}
		if step == 1 &&
			idx > 1 && nums[idx-1] < target &&
			idx < len(nums)-1 && nums[idx+1] > target {
			break
		}
	}
	if nums[idx] == target {
		return idx
	}
	return -1
}
