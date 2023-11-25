package leetcode

func minSubArrayLen(target int, nums []int) int {
	l := len(nums)
	min := 0
	sum := nums[0]
	i, j := 0, 0
	for {
		if sum < target {
			j++
			if j == l {
				break
			}
			sum += nums[j]
			continue
		}
		if sum >= target {
			// 满足条件
			length := j - i + 1
			if min == 0 || length < min {
				min = length
			}
			sum -= nums[i]
			i++
			if i == l {
				break
			}
			continue
		}
	}
	return min
}
