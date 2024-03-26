package leetcode

func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	preMax, maxRes := nums[0], nums[0]
	preMin, minRes := nums[0], nums[0]
	sum := nums[0]
	for i := 1; i < n; i++ {
		preMax = max(preMax+nums[i], nums[i])
		maxRes = max(maxRes, preMax)
		preMin = min(preMin+nums[i], nums[i])
		minRes = min(minRes, preMin)
		sum += nums[i]
	}
	if maxRes < 0 {
		return maxRes
	} else {
		return max(maxRes, sum-minRes)
	}
}
