package leetcode

func lengthOfLIS(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	dp := make([]int, l)

	for i:=0; i<l; i++ {
		dp[i] = 1
		for j:=0; j<i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max4s(dp[i], dp[j]+1)
			}
		}
	}

	return max4s(dp...)
}
