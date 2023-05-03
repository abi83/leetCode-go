package LeetCode75

type leftRightSum struct {
	left  int
	right int
}

func PivotIndex(nums []int) int {
	numsLen := len(nums)
	var caches = make([]leftRightSum, numsLen)

	for i := 1; i < numsLen; i++ {
		caches[i] = leftRightSum{
			left:  caches[i-1].left + nums[i-1],
			right: caches[i].right,
		}
		caches[numsLen-1-i] = leftRightSum{
			left:  caches[numsLen-1-i].left,
			right: caches[numsLen-i].right + nums[numsLen-i],
		}
	}
	for index, value := range caches {
		if value.left == value.right {
			return index
		}
	}
	return -1
}
