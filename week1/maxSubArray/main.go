package maxSubArray

import "math"

func maxSubArray(nums []int) int {
	var maxFirst, maxEnd float64
	maxEnd = float64(nums[0])
	maxFirst = float64(nums[0])
	for i := 1; i < len(nums); i++ {
		maxEnd = math.Max(maxEnd + float64(nums[i]), float64(nums[i]))
		maxFirst = math.Max(maxFirst, maxEnd)
	}

	return int(maxFirst)
}