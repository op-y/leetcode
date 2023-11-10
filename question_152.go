package leetcode

func maxProduct(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}

	res := nums[0]
	lastMax := nums[0]
	lastMin := nums[0]

	for i := 1; i < l; i++ {
		tmpMax := max(nums[i], max(lastMax*nums[i], lastMin*nums[i]))
		tmpMin := min(nums[i], min(lastMin*nums[i], lastMax*nums[i]))
		lastMax = tmpMax
		lastMin = tmpMin
		if lastMax > res {
			res = lastMax
		}
	}

	return res
}
