package leetcode

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	volume := 0
	left := 0
	right := len(height) - 1
	max := Volume(height, left, right)

	// 双指针 移动小的一端边界
	for left < right {
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		if left == right {
			break
		}
		volume = Volume(height, left, right)
		if volume > max {
			max = volume
		}
	}
	return max
}

func Volume(height []int, left, right int) int {
	if height[left] < height[right] {
		return height[left] * (right - left)
	} else {
		return height[right] * (right - left)
	}
}
