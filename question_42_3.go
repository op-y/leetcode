package leetcode

func trap3(height []int) int {
	left, right := 0, n - 1
	leftMax, rightMax := 0, 0

	ans := 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightmax = max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	} 
    return ans
}
